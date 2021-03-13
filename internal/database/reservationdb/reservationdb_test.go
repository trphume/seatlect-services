package reservationdb

import (
	"context"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"testing"
	"time"
)

const (
	jakeID     = "5facaf3bd646b77f40481343"
	brightioID = "5facafef6b28446f285d7ae4"
)

type ReservationSuite struct {
	suite.Suite
	ReservationDB *ReservationDB
}

func (r *ReservationSuite) SetupSuite() {
	// Create MongoDB client and verify connection
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		r.T().Fatal("Mongo connection URI is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		r.T().Fatal("Could create a mongo client")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		r.T().Fatal("Could not connect to mongodb")
	}

	db := client.Database("test")

	// Attach CustomerDB type to Suite
	r.ReservationDB = &ReservationDB{
		ResCol: db.Collection("reservation"),
	}
}

func TestReservationSuite(t *testing.T) {
	suite.Run(t, new(ReservationSuite))
}
