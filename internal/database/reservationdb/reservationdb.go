package reservationdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ReservationDB struct {
	ResCol *mongo.Collection
	BusCol *mongo.Collection
}

func (r *ReservationDB) ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, commonErr.INVALID
	}

	var reservations *mongo.Cursor
	if start.IsZero() || end.IsZero() {
		reservations, err = r.ResCol.Find(ctx, bson.M{"businessId": pId})
	} else {
		reservations, err = r.ResCol.Find(
			ctx,
			bson.D{
				{"businessId", pId},
				{"start", bson.D{
					{"$gte", start},
				}},
				{"end", bson.D{
					{"$lte", end},
				}},
			},
		)
	}

	if err != nil {
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Reservation
	if err = reservations.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (r *ReservationDB) CreateReservation(ctx context.Context, reservation typedb.Reservation) error {
	// get business object first - this is needed to retrieve the placement schema and current display image
	businessResult := r.BusCol.FindOne(ctx, bson.M{"_id": reservation.BusinessId})
	if businessResult.Err() != nil {
		if businessResult.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	var business typedb.Business
	if err := businessResult.Decode(&business); err != nil {
		return commonErr.INTERNAL
	}

	// populate the rest of the reservation data
	pmt := business.Placement
	reservation.Placement = typedb.ReservationPlacement{
		Width:  pmt.Width,
		Height: pmt.Height,
		Seats:  toReservationSeats(pmt.Seats),
	}

	// create in database
	if _, err := r.ResCol.InsertOne(ctx, reservation); err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

func (r *ReservationDB) ReserveSeats(ctx context.Context, id string, user string, seats []string) (typedb.Order, error) {
	panic("implement me")
}

// Parsing function
func toReservationSeats(seats []typedb.Seat) []typedb.ReservationSeat {
	res := make([]typedb.ReservationSeat, len(seats))
	for i, s := range seats {
		res[i] = typedb.ReservationSeat{
			Name:     s.Name,
			Floor:    s.Floor,
			Type:     s.Type,
			Space:    s.Space,
			User:     nil,
			Status:   "EMPTY",
			X:        s.X,
			Y:        s.Y,
			Width:    s.Width,
			Height:   s.Height,
			Rotation: s.Rotation,
		}
	}

	return res
}
