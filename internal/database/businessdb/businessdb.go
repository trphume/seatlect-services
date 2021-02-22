package businessdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BusinessDB struct {
	BusCol *mongo.Collection
}

// TODO: proper search functionality
func (b *BusinessDB) ListBusiness(ctx context.Context, searchParams typedb.ListBusinessParams) ([]typedb.Business, error) {
	// Construcut params
	limit := new(int64)
	*limit = int64(searchParams.Limit)

	// TODO: construct sorting option

	businesses, err := b.BusCol.Find(ctx, bson.M{}, &options.FindOptions{
		Limit: limit,
	})

	if err != nil {
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Business
	if err = businesses.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (b *BusinessDB) ListBusinessByIds(ctx context.Context, ids []string) ([]typedb.Business, error) {
	objIds := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objIds[i], _ = primitive.ObjectIDFromHex(id)
	}

	businesses, err := b.BusCol.Find(
		ctx,
		bson.D{
			{"_id",
				bson.D{
					{"$in", objIds},
				},
			},
		},
	)

	if err != nil {
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Business
	if err = businesses.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}
