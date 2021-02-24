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

func TestBusinessSuite(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
