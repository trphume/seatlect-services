package businessdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/mongo"
)

type BusinessDB struct {
	Mongo *mongo.Collection
	Redis *redis.Client
}

func (b *BusinessDB) ListBusiness(ctx context.Context, searchParams typedb.ListBusinessParams) ([]typedb.Business, error) {
	panic("implement me")
}

func (b *BusinessDB) ListBusinessByIds(ctx context.Context, ids []string) ([]typedb.Business, error) {
	panic("implement me")
}
