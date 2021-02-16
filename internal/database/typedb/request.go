package typedb

import "time"

type Request struct {
	Id           string    `bson:"_id"`
	BusinessName string    `bson:"businessName"`
	Type         string    `bson:"type"`
	Tags         []string  `bson:"tags"`
	Description  string    `bson:"description"`
	Location     Location  `bson:"location"`
	Address      string    `bson:"address"`
	CreatedAt    time.Time `bson:"createdAt"`
}
