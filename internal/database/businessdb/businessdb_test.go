package businessdb

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
	brightioID   = "5facafef6b28446f285d7ae4"
	beerBurgerId = "5facaff31c6d49b2c7256bf3"
)

type BusinessSuite struct {
	suite.Suite
	BusinessDB *BusinessDB
}

func (b *BusinessSuite) SetupSuite() {
	// Create MongoDB client and verify connection
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		b.T().Fatal("Mongo connection URI is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		b.T().Fatal("Could create a mongo client")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		b.T().Fatal("Could not connect to mongodb")
	}

	db := client.Database("test")

	// Attach CustomerDB type to Suite
	b.BusinessDB = &BusinessDB{
		BusCol: db.Collection("business"),
	}
}

func (b *BusinessSuite) TestListBusinessByIds() {
	tests := []struct {
		in  []string
		out int
		err error
	}{
		{in: []string{brightioID}, out: 1, err: nil},
		{in: []string{brightioID, beerBurgerId}, out: 2, err: nil},
		{in: []string{brightioID, "somerandomid"}, out: 1, err: nil},
		{in: []string{"somerandomid"}, out: 0, err: nil},
		{in: []string{}, out: 0, err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		out, err := b.BusinessDB.ListBusinessByIds(ctx, tt.in)

		b.Assert().Equal(tt.err, err)
		b.Assert().Equal(tt.out, len(out))
	}
}

func TestBusinessSuite(t *testing.T) {
	suite.Run(t, new(BusinessSuite))
}
