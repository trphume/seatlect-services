package typedb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reservation struct {
	Id         primitive.ObjectID   `bson:"_id"`
	BusinessId primitive.ObjectID   `bson:"businessId"`
	Name       string               `bson:"name"`
	Start      time.Time            `bson:"start"`
	End        time.Time            `bson:"end"`
	Placement  ReservationPlacement `bson:"placement"`
	Image      string               `bson:"image"`
	Location   Location             `bson:"location"`
	Type       string               `bson:"type"`
	Status     int                  `bson:"status"`
}

type ReservationSeat struct {
	Name     string             `bson:"name"`
	Floor    int                `bson:"floor"`
	Type     string             `bson:"type"`
	Space    int                `bson:"space"`
	User     primitive.ObjectID `bson:"user"`
	Username string             `bson:"username"`
	Status   string             `bson:"status"`
	X        float64            `bson:"x"`
	Y        float64            `bson:"y"`
	Width    float64            `bson:"width"`
	Height   float64            `bson:"height"`
	Rotation float64            `bson:"rotation"`
}

type ReservationPlacement struct {
	Width  int               `bson:"width"`
	Height int               `bson:"height"`
	Seats  []ReservationSeat `bson:"seats"`
}

type SearchReservationParams struct {
	Name     string
	Type     string
	Location Location
	Start    time.Time
	End      time.Time
}
