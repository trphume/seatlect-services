package reservationdb

import (
	"context"
	"fmt"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
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
		reservations, err = r.ResCol.Find(ctx, bson.M{"businessId": pId, "status": 1, "start": bson.M{"$gte": time.Now()}})
	} else {
		reservations, err = r.ResCol.Find(
			ctx,
			bson.D{
				{"businessId", pId},
				{"status", 1},
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
func (r *ReservationDB) SearchReservation(ctx context.Context, searchParams typedb.SearchReservationParams) ([]typedb.Reservation, error) {
	// Construct query
	query := bson.D{
		{"location", bson.D{
			{"$geoWithin", bson.M{"$centerSphere": bson.A{searchParams.Location.Coordinates, 0.00156786503}}},
		}},
		{"type", searchParams.Type},
		{"status", 1},
		{"start", bson.D{
			{"$gte", searchParams.Start},
		}},
		{"end", bson.D{
			{"$lte", searchParams.End},
		}},
	}

	if searchParams.Name != "" {
		query = bson.D{
			{"$text", bson.M{"$search": searchParams.Name}},
			{"location", bson.D{
				{"$geoWithin", bson.M{"$centerSphere": bson.A{searchParams.Location.Coordinates, 0.00156786503}}},
			}},
			{"type", searchParams.Type},
			{"status", 1},
			{"start", bson.D{
				{"$gte", searchParams.Start},
			}},
			{"end", bson.D{
				{"$lte", searchParams.End},
			}},
		}
	}

	// Database call
	reservations, err := r.ResCol.Find(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Reservation
	if err = reservations.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (r *ReservationDB) CreateReservation(ctx context.Context, reservation *typedb.Reservation) error {
	// check that there is no overlapping reservation time first
	overlapping := r.ResCol.FindOne(
		ctx,
		bson.M{
			"businessId": reservation.BusinessId,
			"status":     1,
			"$or": bson.A{
				bson.M{"start": bson.M{
					"$gte": reservation.Start,
					"$lte": reservation.End,
				}},
				bson.M{
					"end": bson.M{
						"$gte": reservation.Start,
						"$lte": reservation.End,
					},
				},
				bson.M{
					"$and": bson.A{
						bson.M{"start": bson.M{"$lte": reservation.Start}},
						bson.M{"end": bson.M{"$gte": reservation.End}},
					},
				},
			},
		},
		options.FindOne().SetProjection(bson.M{"_id": 1}),
	)

	if overlapping.Err() != mongo.ErrNoDocuments {
		return commonErr.CONFLICT
	}

	// get business object first - this is needed to retrieve the placement schema and current display image
	businessResult := r.BusCol.FindOne(
		ctx,
		bson.M{"_id": reservation.BusinessId},
		options.FindOne().SetProjection(bson.M{"placement": 1, "businessName": 1, "location": 1, "type": 1, "displayImage": 1}),
	)

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

	// Construct any other missing field
	// if a reservation name was not supplied
	if reservation.Name == "" {
		reservation.Name = business.BusinessName
	}

	reservation.Location = business.Location
	reservation.Type = business.Type
	reservation.Image = business.DisplayImage

	// create in database
	if _, err := r.ResCol.InsertOne(ctx, reservation); err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

func (r *ReservationDB) ReserveSeats(ctx context.Context, id string, user string, username string, seats []string) (*typedb.Order, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, commonErr.INVALID
	}

	uId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return nil, commonErr.INVALID
	}

	if len(seats) == 0 {
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
			if s.Name == n && s.Status == "AVAILABLE" {
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
					"placement.seats.$[elem].user":     uId,
					"placement.seats.$[elem].username": username,
					"placement.seats.$[elem].status":   "TAKEN",
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
		Status:        "AVAILABLE",
		Image:         resv.Image,
		ExtraSpace:    0,
		Name:          resv.Name,
	}

	if _, err := r.OrdCol.InsertOne(ctx, order); err != nil {
		return nil, commonErr.INTERNAL
	}

	return &order, nil
}

func (r *ReservationDB) GetReservationById(ctx context.Context, businessId string, reservationId string) (*typedb.Reservation, error) {
	rId, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return nil, commonErr.INVALID
	}

	var bId primitive.ObjectID
	if businessId != "" {
		bId, err = primitive.ObjectIDFromHex(businessId)
		if err != nil {
			return nil, commonErr.INVALID
		}
	}

	var query bson.M
	if businessId != "" {
		query = bson.M{"_id": rId, "businessId": bId}
	} else {
		query = bson.M{"_id": rId}
	}

	res := r.ResCol.FindOne(ctx, query)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, commonErr.NOTFOUND
		}

		return nil, commonErr.INTERNAL
	}

	var reservation typedb.Reservation
	if err := res.Decode(&reservation); err != nil {
		return nil, commonErr.INTERNAL
	}

	return &reservation, nil
}

func (r *ReservationDB) UpdateReservationStatus(ctx context.Context, reservationId string) ([]primitive.ObjectID, error) {
	rId, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return nil, commonErr.INVALID
	}

	// update query
	reservation := r.ResCol.FindOneAndUpdate(
		ctx,
		bson.M{"_id": rId},
		bson.M{"$set": bson.M{"status": 0}},
		options.FindOneAndUpdate().SetProjection(bson.M{"placement": 1}),
	)

	if reservation.Err() != nil {
		if reservation.Err() == mongo.ErrNoDocuments {
			return nil, commonErr.NOTFOUND
		}

		return nil, commonErr.INVALID
	}

	// parse reservation.placement to get list of user id
	var resv typedb.Reservation
	if err := reservation.Decode(&resv); err != nil {
		return nil, commonErr.INVALID
	}

	tmp := make(map[primitive.ObjectID]bool)
	for _, s := range resv.Placement.Seats {
		tmp[s.User] = true
	}

	res := make([]primitive.ObjectID, len(tmp))
	i := 0
	for k := range tmp {
		res[i] = k
		i++
	}

	return res, nil
}

func (r *ReservationDB) UpdateOrderStatus(ctx context.Context, reservationId string) error {
	rId, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return commonErr.INVALID
	}

	_, err = r.OrdCol.UpdateMany(ctx, bson.M{"reservationId": rId}, bson.M{"$set": bson.M{"status": "CANCELLED"}})
	if err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

// Parsing function
func toReservationSeats(seats []typedb.Seat) []typedb.ReservationSeat {
	res := make([]typedb.ReservationSeat, len(seats))
	for i, s := range seats {
		var status string
		if strings.Contains(s.Type, "TABLE") {
			status = "AVAILABLE"
		} else {
			status = ""
		}

		res[i] = typedb.ReservationSeat{
			Name:     s.Name,
			Floor:    s.Floor,
			Type:     s.Type,
			Space:    s.Space,
			User:     primitive.ObjectID{},
			Status:   status,
			X:        s.X,
			Y:        s.Y,
			Width:    s.Width,
			Height:   s.Height,
			Rotation: s.Rotation,
		}
	}

	return res
}
