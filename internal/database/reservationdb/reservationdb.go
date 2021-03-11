package reservationdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ReservationDB struct {
	ResCol *mongo.Collection
}

func (r *ReservationDB) ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error) {
	panic("implement me")
}
