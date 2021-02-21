package customerdb

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
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

type CustomerSuite struct {
	suite.Suite
	CustomerDB *CustomerDB
}

func (c *CustomerSuite) SetupSuite() {
	// Create MongoDB client and verify connection
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
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

	db := client.Database("test")

	// Attach CustomerDB type to Suite
	c.CustomerDB = &CustomerDB{
		CusCol: db.Collection("customer"),
		BusCol: db.Collection("business"),
	}
}

func (c *CustomerSuite) TestAuthenticateCustomer() {
	tests := []struct {
		in  *typedb.Customer
		out string
		err error
	}{
		{in: &typedb.Customer{Username: "Jake", Password: "ExamplePassword"}, out: jakeID, err: nil},
		{in: &typedb.Customer{Username: "DoesNotExist", Password: "DoesNotMatter"}, out: "", err: commonErr.NOTFOUND},
		{in: &typedb.Customer{Username: "Jake", Password: "WrongPassword"}, out: "", err: commonErr.NOTFOUND},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		out, err := c.CustomerDB.AuthenticateCustomer(ctx, tt.in)

		c.Assert().Equal(tt.out, out)
		c.Assert().Equal(tt.err, err)
	}
}

func (c *CustomerSuite) TestCreateCustomer() {
	tests := []struct {
		in  *typedb.Customer
		err error
	}{
		{in: &typedb.Customer{Username: "Jake", Password: "ExamplePassword", Dob: time.Unix(1613905601, 0), Favorite: make([]string, 0)}, err: commonErr.INTERNAL},
		{in: &typedb.Customer{Username: "Tata", Password: "DoesNotMatter", Dob: time.Unix(1613905601, 0), Favorite: make([]string, 0)}, err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err := c.CustomerDB.CreateCustomer(ctx, tt.in)

		c.Assert().Equal(tt.err, err)
	}
}

func (c *CustomerSuite) TestAppendFavorite() {
	tests := []struct {
		in  []string
		err error
	}{
		{in: []string{jakeID, brightioID}, err: nil},
		{in: []string{jakeID, "DoesNotExist"}, err: commonErr.NOTFOUND},
		{in: []string{jakeID, brightioID}, err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := c.CustomerDB.AppendFavorite(ctx, tt.in[0], tt.in[1])

		c.Assert().Equal(tt.err, err)
	}
}

func (c *CustomerSuite) TestRemoveFavorite() {}

func TestCustomerSuite(t *testing.T) {
	suite.Run(t, new(CustomerSuite))
}
