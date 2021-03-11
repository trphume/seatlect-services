package reservationdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ReservationDB struct {
	ResCol *mongo.Collection
}

func (r *ReservationDB) ListReservation(ctx context.Context, id string, start time.Time, end time.Time) {
	panic("implement me")
}
