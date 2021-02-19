package typedb

import "time"

type Customer struct {
	Id       string    `bson:"_id"`
	Username string    `bson:"username"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
	Dob      time.Time `bson:"dob"`
	Favorite []string  `bson:"favorite"`
	Verified bool      `bson:"verified"`
}
