package typedb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Request struct {
	Id           primitive.ObjectID `bson:"_id"`
	BusinessName string             `bson:"businessName"`
	Type         string             `bson:"type"`
	Tags         []string           `bson:"tags"`
	Description  string             `bson:"description"`
	Location     Location           `bson:"location"`
	Address      string             `bson:"address"`
	CreatedAt    time.Time          `bson:"createdAt"`
}
