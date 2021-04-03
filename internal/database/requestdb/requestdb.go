package requestdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RequestDB struct {
	ReqCol *mongo.Collection
	BusCol *mongo.Collection
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
	if err = req.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (r *RequestDB) ApproveRequest(ctx context.Context, id string) error {
	rId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	// Find the request and delete
	tmp := r.ReqCol.FindOneAndDelete(ctx, bson.M{"_id": rId})
	if tmp.Err() != nil {
		if tmp.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	var req typedb.Request
	if err = tmp.Decode(&req); err != nil {
		return commonErr.INTERNAL
	}

	// Update information in business collection
	_, err = r.BusCol.UpdateOne(
		ctx,
		bson.M{"_id": req.Id},
		bson.D{
			{"$set",
				bson.D{
					{"businessName", req.BusinessName},
					{"type", req.Type},
					{"tags", req.Tags},
					{"description", req.Description},
					{"location", req.Location},
					{"address", req.Address},
				},
			},
		},
	)

	if err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

func (r *RequestDB) GetRequestById(ctx context.Context, request *typedb.Request) error {
	res := r.ReqCol.FindOne(ctx, bson.M{"_id": request.Id})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	if err := res.Decode(request); err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

func (r *RequestDB) CreateRequest(ctx context.Context, request *typedb.Request) error {
	// remove any old request first
	_, err := r.ReqCol.DeleteOne(ctx, bson.M{"_id": request.Id})
	if err != nil {
		return commonErr.INTERNAL
	}

	// insert new request
	_, err = r.ReqCol.InsertOne(ctx, request)
	if err != nil {
		return commonErr.INTERNAL
	}

	return nil
}

func (r *RequestDB) DeleteRequest(ctx context.Context, id string) error {
	rId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	res, err := r.ReqCol.DeleteOne(ctx, bson.M{"_id": rId})
	if err != nil {
		return commonErr.INTERNAL
	}

	if res.DeletedCount == 0 {
		return commonErr.NOTFOUND
	}

	return nil
}

// helper function
func createBool(b bool) *bool {
	return &b
}
