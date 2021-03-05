package requestdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestDB struct {
	ReqCol *mongo.Collection
}

func (r *RequestDB) ListRequest(ctx context.Context, page int, requests []typedb.Request) (int, error) {
	panic("implement me")
}

func (r *RequestDB) ApproveRequest(ctX context.Context, id string) error {
	panic("implement me")
}

func (r *RequestDB) GetRequestById(ctx context.Context, request *typedb.Request) error {
	panic("implement me")
}

func (r *RequestDB) CreateRequest(ctx context.Context, request *typedb.Request) error {
	panic("implement me")
}
