package admindb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminDB struct {
	AdminCol *mongo.Collection
}

func (a *AdminDB) AuthenticateAdmin(ctx context.Context, username string, password string) (string, error) {
	panic("implement me")
}
