package placementdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlacementDB struct {
	BusCol *mongo.Collection
}

func (p *PlacementDB) GetPlacement(ctx context.Context, id string) (typedb.Placement, error) {
	panic("implement me")
}

func (p *PlacementDB) UpdatePlacement(ctx context.Context, id string, placement typedb.Placement) error {
	panic("implement me")
}
