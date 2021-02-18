package customerdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerDB struct {
	Mongo *mongo.Collection
	Redis *redis.Client
}

func (c *CustomerDB) AuthenticateCustomer(ctx context.Context, customer *typedb.Customer, password string) (string, error) {
	panic("implement me")
}

func (c *CustomerDB) CreateCustomer(ctx context.Context, customer *typedb.Customer, password string) (string, error) {
	panic("implement me")
}

func (c *CustomerDB) AppendFavorite(ctx context.Context, businessId string) error {
	panic("implement me")
}

func (c *CustomerDB) RemoveFavorite(ctx context.Context, businessId string) error {
	panic("implement me")
}
