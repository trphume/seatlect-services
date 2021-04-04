package placementdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlacementDB struct {
	BusCol *mongo.Collection
}

func (p *PlacementDB) GetPlacement(ctx context.Context, id string) (*typedb.Placement, error) {
	// Slightly modified version from GetBusinessById from businessdb
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, commonErr.INVALID
	}

	business := p.BusCol.FindOne(ctx, bson.M{"_id": pId})
	if business.Err() != nil {
		if business.Err() == mongo.ErrNoDocuments {
			return nil, commonErr.NOTFOUND
		}

		return nil, commonErr.INTERNAL
	}

	var res typedb.Business
	if err = business.Decode(&res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return &res.Placement, nil
}

func (p *PlacementDB) UpdatePlacement(ctx context.Context, id string, placement typedb.Placement) error {
	// Slightly modified version from UpdateBusinessById from businessdb
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	res, err := p.BusCol.UpdateOne(
		ctx,
		bson.M{"_id": pId},
		bson.D{
			{"$set",
				bson.D{
					{"placement", placement},
				},
			},
		},
	)

	if err != nil {
		return commonErr.INTERNAL
	}

	if res.ModifiedCount == 0 {
		return commonErr.NOTFOUND
	}

	return nil
}
