package typedb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	Id            primitive.ObjectID `bson:"_id"`
	ReservationId primitive.ObjectID `bson:"reservationId"`
	CustomerId    primitive.ObjectID `bson:"customerId"`
	BusinessId    primitive.ObjectID `bson:"businessId"`
	Start         time.Time          `bson:"start"`
	End           time.Time          `bson:"end"`
	Seats         []Seat             `bson:"seats"`
	Status        string             `bson:"status"`
	Image         string             `bson:"image"`
	ExtraSpace    int                `bson:"extraSpace"`
	Name          string             `bson:"name"`
}
