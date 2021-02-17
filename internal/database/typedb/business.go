package typedb

import (
	"time"
)

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

type ListBusinessParams struct {
	Limit      int32
	Sort       BusinessSort
	Name       string
	Type       string
	Tags       []string
	Location   Location
	StartPrice int32
	EndPrice   int32
	StartDate  time.Time
	EndDate    time.Time
}

// Helper definition
type BusinessSort int32

const (
	Sort_NAME_ASC      BusinessSort = 0
	Sort_NAME_DESC     BusinessSort = 1
	Sort_LOCATION_ASC  BusinessSort = 2
	Sort_LOCATION_DESC BusinessSort = 3
	Sort_PRICE_ASC     BusinessSort = 4
	Sort_PRICE_DESC    BusinessSort = 5
	Sort_DATE_ASC      BusinessSort = 6
	Sort_DATE_DESC     BusinessSort = 7
)
