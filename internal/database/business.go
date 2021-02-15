package database

type Business struct {
	Id           string      `bson:"_id"`
	Username     string      `bson:"username"`
	Email        string      `bson:"email"`
	BusinessName string      `bson:"businessName"`
	Type         string      `bson:"type"`
	Tags         []string    `bson:"tags"`
	Description  string      `bson:"description"`
	Location     Location    `bson:"location"`
	Address      string      `bson:"address"`
	DisplayImage string      `bson:"displayImage"`
	Images       []string    `bson:"images"`
	Placement    Seat        `bson:"placement"`
	Menu         []MenuItems `bson:"menu"`
	Status       int         `bson:"status"`
	Verified     bool        `bson:"verified"`
}
