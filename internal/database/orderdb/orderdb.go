package orderdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderDB struct {
	Mongo *mongo.Collection
}

func (o *OrderDB) ListOrderByCustomer(ctx context.Context, CustomerId string, limit int32, page int32) ([]typedb.Order, error) {
	panic("implement me")
}
