package typedb

type Admin struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
}
