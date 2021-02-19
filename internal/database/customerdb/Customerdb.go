package customerdb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type CustomerDB struct {
	Mongo *mongo.Collection
}

func (c *CustomerDB) AuthenticateCustomer(ctx context.Context, customer *typedb.Customer) (string, error) {
	res := c.Mongo.FindOne(ctx, bson.M{"_id": customer.Id})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", commonErr.NOTFOUND
		}

		return "", commonErr.INTERNAL
	}

	pw := customer.Password
	if err := res.Decode(customer); err != nil {
		return "", commonErr.INTERNAL
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(pw)); err != nil {
		return "", commonErr.NOTFOUND
	}

	return customer.Id, nil
}

func (c *CustomerDB) CreateCustomer(ctx context.Context, customer *typedb.Customer) (string, error) {
	panic("implement me")
}

func (c *CustomerDB) AppendFavorite(ctx context.Context, customerId string, businessId string) error {
	panic("implement me")
}

func (c *CustomerDB) RemoveFavorite(ctx context.Context, customerId string, businessId string) error {
	panic("implement me")
}
