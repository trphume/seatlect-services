package typedb

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Seat struct {
	Name     string  `bson:"name"`
	Floor    int     `bson:"floor"`
	Type     string  `bson:"type"`
	Space    int     `bson:"space"`
	X        float64 `bson:"x"`
	Y        float64 `bson:"y"`
	Width    float64 `bson:"width"`
	Height   float64 `bson:"height"`
	Rotation float64 `bson:"rotation"`
}

type Placement struct {
	Width  int    `bson:"width"`
	Height int    `bson:"height"`
	Seats  []Seat `bson:"seats"`
}

type MenuItems struct {
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Image       string  `bson:"image"`
	Price       float64 `bson:"price"`
}
