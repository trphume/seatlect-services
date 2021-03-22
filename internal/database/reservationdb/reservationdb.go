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
