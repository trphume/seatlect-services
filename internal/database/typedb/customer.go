package typedb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Customer struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Dob      time.Time          `bson:"dob"`
	Favorite []string           `bson:"favorite"`
	Verified bool               `bson:"verified"`
}
