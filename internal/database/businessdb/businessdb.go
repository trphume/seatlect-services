package businessdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

func (b *BusinessDB) AuthenticateBusiness(ctx context.Context, business *typedb.Business) (string, error) {
	res := b.BusCol.FindOne(ctx, bson.M{"username": business.Username})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", commonErr.NOTFOUND
		}

		return "", commonErr.INTERNAL
	}

	pw := business.Password
	if err := res.Decode(business); err != nil {
		return "", commonErr.INTERNAL
	}

	if err := bcrypt.CompareHashAndPassword([]byte(business.Password), []byte(pw)); err != nil {
		return "", commonErr.NOTFOUND
	}

	return business.Id.Hex(), nil
}

func (b *BusinessDB) CreateBusiness(ctx context.Context, business *typedb.Business) error {
	panic("implement me")
}

func (b *BusinessDB) SimpleListBusiness(ctx context.Context, status int, page int, business []typedb.Business) (int, error) {
	panic("implement me")
}

func (b *BusinessDB) GetBusinessById(ctx context.Context, id string) (typedb.Business, error) {
	panic("implement me")
}

func (b *BusinessDB) UpdateBusinessById(ctx context.Context, business typedb.Business) error {
	panic("implement me")
}

func (b *BusinessDB) UpdateBusinessDIById(ctx context.Context, id string, image string) (string, error) {
	panic("implement me")
}

func (b *BusinessDB) AppendBusinessImage(ctx context.Context, id string, image string) error {
	panic("implement me")
}

func (b *BusinessDB) RemoveBusinessImage(ctx context.Context, id string, pos int) error {
	panic("implement me")
}

func (b *BusinessDB) ListMenuItem(ctx context.Context, id string) ([]typedb.MenuItems, error) {
	panic("implement me")
}

func (b *BusinessDB) AppendMenuItem(ctx context.Context, id string, item typedb.MenuItems) (string, error) {
	panic("implement me")
}

func (b *BusinessDB) RemoveMenuItem(ctx context.Context, id string, name string) error {
	panic("implement me")
}

func (b *BusinessDB) UpdateBusinessStatus(ctx context.Context, id string, status int) error {
	panic("implement me")
}
