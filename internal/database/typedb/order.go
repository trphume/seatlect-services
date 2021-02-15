package typedb

import "time"

type Order struct {
	Id            string    `bson:"_id"`
	ReservationId string    `bson:"reservationId"`
	CustomerId    string    `bson:"customerId"`
	BusinessId    string    `bson:"businessId"`
	Start         time.Time `bson:"start"`
	End           time.Time `bson:"end"`
	Seats         []Seat    `bson:"seats"`
	Status        string    `bson:"status"`
	Image         string    `bson:"image"`
	ExtraSpace    int       `bson:"extraSpace"`
}
