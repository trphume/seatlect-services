package orderdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderDB struct {
	OrdCol *mongo.Collection
	ResCol *mongo.Collection
}

func (o *OrderDB) ListOrderByCustomer(ctx context.Context, customerId string, limit int32, page int32) ([]typedb.Order, error) {
	pCustomerId, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return nil, commonErr.INVALID
	}

	// construct params
	pLimit := new(int64)
	*pLimit = int64(limit)

	pSkip := new(int64)
	*pSkip = int64(limit * (page - 1))

	orders, err := o.OrdCol.Find(
		ctx,
		bson.D{{"customerId", pCustomerId}},
		&options.FindOptions{Limit: pLimit, Skip: pSkip, Sort: bson.M{"start": 1}},
	)

	if err != nil {
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Order
	if err = orders.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (o *OrderDB) CancelOrder(ctx context.Context, id string) error {
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	// find and delete order
	ordRes := o.OrdCol.FindOneAndDelete(
		ctx,
		bson.M{"_id": oId},
		options.FindOneAndDelete().SetProjection(bson.M{"reservationId": 1, "seats": 1}),
	)

	if ordRes.Err() != nil {
		if ordRes.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	var ord typedb.Order
	if err = ordRes.Decode(&ord); err != nil {
		return commonErr.INTERNAL
	}

	seats := make([]string, len(ord.Seats))
	for i, seat := range ord.Seats {
		seats[i] = seat.Name
	}

	// update status of reservation
	if _, err := o.ResCol.UpdateOne(
		ctx,
		bson.M{"_id": ord.ReservationId},
		bson.D{
			{
				"$set",
				bson.M{
					"placement.seats.$[elem].user":     nil,
					"placement.seats.$[elem].username": "",
					"placement.seats.$[elem].status":   "AVAILABLE",
				},
			},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{Filters: []interface{}{
			bson.M{"elem.name": bson.M{"$in": seats}},
		}}),
	); err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

func (o *OrderDB) GetOrderWithReservationId(ctx context.Context, orderId string, reservationId string) (*typedb.Order, error) {
	// parsing id
	oId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, commonErr.INVALID
	}

	rId, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return nil, commonErr.INVALID
	}

	// query db
	res := o.OrdCol.FindOne(
		ctx,
		bson.M{"_id": oId, "reservationId": rId},
		options.FindOne().SetProjection(bson.M{"seats": 1}),
	)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, commonErr.NOTFOUND
		}

		return nil, commonErr.INTERNAL
	}

	// deocde result
	var order typedb.Order
	if err := res.Decode(&order); err != nil {
		return nil, commonErr.INTERNAL
	}

	return &order, nil
}

func (o *OrderDB) UpdateOrderStatus(ctx context.Context, orderId string, status string) error {
	oId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return commonErr.INVALID
	}

	res, err := o.OrdCol.UpdateOne(ctx, bson.M{"_id": oId}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return commonErr.INTERNAL
	}

	if res.MatchedCount == 0 {
		return commonErr.NOTFOUND
	}

	return nil
}
