package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Seat struct {
	Name     string  `bson:"Name"`
	Floor    int     `bson:"floor"`
	Type     string  `bson:"type"`
	Space    string  `bson:"space"`
	X        float64 `bson:"x"`
	Y        float64 `bson:"y"`
	Width    float64 `bson:"width"`
	Height   float64 `bson:"height"`
	Rotation float64 `bson:"rotation"`
}

type MenuItems struct {
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Image       string               `bson:"image"`
	Price       primitive.Decimal128 `bson:"price"`
}
