package requestdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RequestDB struct {
	ReqCol *mongo.Collection
}

func (r *RequestDB) ListRequest(ctx context.Context, page int) ([]typedb.Request, error) {
	if page == 0 {
		page = 1
	}

	// Construct params
	limit := int64(10)

	p := int64((page - 1) * 10)

	req, err := r.ReqCol.Find(ctx, bson.M{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &p,
	})

	if err != nil {
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Request
	if err == req.All(ctx, res) {
		return nil, commonErr.INTERNAL
	}

	return res, nil
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
