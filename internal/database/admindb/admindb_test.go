package admindb

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

const admin1ID = "604dfa455226a8714411f33d"

type AdminSuite struct {
	suite.Suite
	AdminDB *AdminDB
}

func (a *AdminSuite) SetupSuite() {
	// Create MongoDB client and verify connection
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		a.T().Fatal("Mongo connection URI is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		a.T().Fatal("Could create a mongo client")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		a.T().Fatal("Could not connect to mongodb")
	}

	db := client.Database("test")

	// Attach CustomerDB type to Suite
	a.AdminDB = &AdminDB{
		AdminCol: db.Collection("admin"),
	}
}

func (a *AdminSuite) TestAuthenticateAdmin() {
	tests := []struct {
		username string
		password string
		out      string
		err      error
	}{
		{username: "admin1", password: "ExamplePassword", out: admin1ID, err: nil},
		{username: "admin2", password: "ExamplePassword", out: "", err: commonErr.NOTFOUND},
		{username: "admin1", password: "wrongpassword", out: "", err: commonErr.NOTFOUND},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		id, err := a.AdminDB.AuthenticateAdmin(ctx, tt.username, tt.password)

		a.Assert().Equal(tt.err, err)
		a.Assert().Equal(tt.out, id)
	}
}

func TestAdminSuite(t *testing.T) {
	suite.Run(t, new(AdminSuite))
}
