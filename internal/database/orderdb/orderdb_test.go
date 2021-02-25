package orderdb

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

type OrderSuite struct {
	suite.Suite
	OrderDB *OrderDB
}

func (b *OrderSuite) SetupSuite() {
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
	b.OrderDB = &OrderDB{
		OrdCol: db.Collection("order"),
	}
}

func (b *OrderSuite) TestListOrderByCustomer() {
	tests := []struct {
		in  string
		out int
		err error
	}{
		{in: jakeID, out: 1, err: nil},
		{in: brightioID, out: 0, err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		out, err := b.OrderDB.ListOrderByCustomer(ctx, tt.in, 10, 1)

		b.Assert().Equal(tt.err, err)
		b.Assert().Equal(tt.out, len(out))
	}
}

func TestBusinessSuite(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
