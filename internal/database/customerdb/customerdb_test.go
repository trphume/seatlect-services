package customerdb

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

type CustomerSuite struct {
	suite.Suite
	CustomerDB *CustomerDB
}

func (c *CustomerSuite) SetupSuite() {
	// Create MongoDB client and verify connection
	mongoURI := os.Getenv("MONGO_URI")
	if c.Assert().Empty(mongoURI) {
		c.T().Fatal("Mongo connection URI is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		c.T().Fatal("Could create a mongo client")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		c.T().Fatal("Could not connect to mongodb")
	}

	db := client.Database("db")

	// Attach CustomerDB type to Suite
	c.CustomerDB = &CustomerDB{
		CusCol: db.Collection("customer"),
		BusCol: db.Collection("business"),
	}
}

func (c *CustomerSuite) TestAuthenticateCustomer() {}

func (c *CustomerSuite) TestCreateCustomer() {}

func (c *CustomerSuite) TestAppendFavorite() {}

func (c *CustomerSuite) TestRemoveFavorite() {}

func TestCustomerSuite(t *testing.T) {
	suite.Run(t, new(CustomerSuite))
}
