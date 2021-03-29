package typedb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Business struct {
	Id           primitive.ObjectID `bson:"_id"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	BusinessName string             `bson:"businessName"`
	Type         string             `bson:"type"`
	Tags         []string           `bson:"tags"`
	Description  string             `bson:"description"`
	Location     Location           `bson:"location"`
	Address      string             `bson:"address"`
	DisplayImage string             `bson:"displayImage"`
	Images       []string           `bson:"images"`
	Placement    Placement          `bson:"placement"`
	Menu         []MenuItems        `bson:"menu"`
	Status       int                `bson:"status"`
	Verified     bool               `bson:"verified"`
}

type ListBusinessParams struct {
	Limit      int32
	Sort       int32
	Name       string
	Type       string
	Tags       []string
	Location   Location
	StartPrice int32
	EndPrice   int32
	StartDate  time.Time
	EndDate    time.Time
}
