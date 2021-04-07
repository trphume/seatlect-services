package reservationdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ReservationDB struct {
	ResCol *mongo.Collection
	BusCol *mongo.Collection
	OrdCol *mongo.Collection
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
			&options.FindOptions{Sort: bson.M{"start": 1}},
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

func (r *ReservationDB) ReserveSeats(ctx context.Context, id string, user string, seats []string) (*typedb.Order, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, commonErr.INVALID
	}

	uId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return nil, commonErr.INVALID
	}

	// check seats availability first
	reservation := r.ResCol.FindOne(ctx, bson.M{"_id": pId})
	if reservation.Err() != nil {
		if reservation.Err() == mongo.ErrNoDocuments {
			return nil, commonErr.NOTFOUND
		}

		return nil, commonErr.INTERNAL
	}

	var resv typedb.Reservation
	if err = reservation.Decode(&resv); err != nil {
		return nil, commonErr.INTERNAL
	}

	count := 0
	typedbSeats := make([]typedb.Seat, len(seats))
	for _, s := range resv.Placement.Seats {
		for _, n := range seats {
			if s.Name == n {
				typedbSeats[count] = typedb.Seat{
					Name:     s.Name,
					Floor:    s.Floor,
					Type:     s.Type,
					Space:    s.Space,
					X:        s.X,
					Y:        s.Y,
					Width:    s.Width,
					Height:   s.Height,
					Rotation: s.Rotation,
				}
				count++
			}
		}
	}

	if count != len(seats) {
		return nil, commonErr.CONFLICT
	}

	// update reservation
	if _, err := r.ResCol.UpdateOne(
		ctx,
		bson.M{"_id": pId},
		bson.D{
			{
				"$set",
				bson.M{
					"placement.seats.$[elem].user":   uId,
					"placement.seats.$[elem].status": "TAKEN",
				},
			},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{Filters: []interface{}{
			bson.M{"elem.name": bson.M{"$in": seats}},
		}}),
	); err != nil {
		return nil, commonErr.INTERNAL
	}

	// create order
	order := typedb.Order{
		Id:            primitive.NewObjectIDFromTimestamp(time.Now()),
		ReservationId: pId,
		CustomerId:    uId,
		BusinessId:    resv.BusinessId,
		Start:         resv.Start,
		End:           resv.End,
		Seats:         typedbSeats,
		Status:        "UNUSED",
		Image:         resv.Image,
		ExtraSpace:    0,
		Name:          resv.Name,
	}

	if _, err := r.OrdCol.InsertOne(ctx, order); err != nil {
		return nil, commonErr.INTERNAL
	}

	return &order, nil
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
			User:     primitive.ObjectID{},
			Status:   "AVAILABLE",
			X:        s.X,
			Y:        s.Y,
			Width:    s.Width,
			Height:   s.Height,
			Rotation: s.Rotation,
		}
	}

	return res
}
