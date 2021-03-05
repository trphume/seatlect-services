package admindb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AdminDB struct {
	AdminCol *mongo.Collection
}

func (a *AdminDB) AuthenticateAdmin(ctx context.Context, username string, password string) (string, error) {
	res := a.AdminCol.FindOne(ctx, bson.M{"username": username})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", commonErr.NOTFOUND
		}

		return "", commonErr.INTERNAL
	}

	var admin typedb.Admin
	if err := res.Decode(&admin); err != nil {
		return "", commonErr.INTERNAL
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", commonErr.NOTFOUND
	}

	return admin.Id.Hex(), nil
}
