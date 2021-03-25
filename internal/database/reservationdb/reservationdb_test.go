package reservationdb

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"testing"
	"time"
)

const (
	jakeID        = "5facaf3bd646b77f40481343"
	brightioID    = "5facafef6b28446f285d7ae4"
	specialTaleID = "5fcde2ec209efa45620a08b6"
	reservationA  = "6035f3a48d505df0b9d043a3"
	reservationB  = "604c80551714a597557abc2e"
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

func (r *ReservationSuite) TestListReservation() {
	tests := []struct {
		in     string
		lenout int
		idout  string
		err    error
	}{
		{in: brightioID, lenout: 1, idout: reservationA, err: nil},
		{in: specialTaleID, lenout: 1, idout: reservationB, err: nil},
		{in: jakeID, lenout: 0, idout: "", err: nil},
		{in: "randomid", lenout: 0, idout: "", err: commonErr.INVALID},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		out, err := r.ReservationDB.ListReservation(ctx, tt.in, time.Time{}, time.Time{})

		r.Assert().Equal(tt.err, err)
		r.Assert().Equal(tt.lenout, len(out))

		if len(out) != 0 {
			r.Assert().Equal(tt.idout, out[0].Id.Hex())
		}
	}
}

func TestReservationSuite(t *testing.T) {
	suite.Run(t, new(ReservationSuite))
}
