package businessdb

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

func (b *BusinessSuite) TestAuthenticateBusiness() {
	tests := []struct {
		in    *typedb.Business
		outid string
		err   error
	}{
		{in: &typedb.Business{Username: "BeerBurger", Password: "ExamplePassword"}, outid: beerBurgerId, err: nil},
		{in: &typedb.Business{Username: "BeerBurger", Password: "WrongPassword"}, err: commonErr.NOTFOUND},
		{in: &typedb.Business{Username: "RandomID", Password: "ExamplePassword"}, err: commonErr.NOTFOUND},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		id, err := b.BusinessDB.AuthenticateBusiness(ctx, tt.in)

		b.Assert().Equal(tt.err, err)
		b.Assert().Equal(tt.outid, id)
	}
}

func (b *BusinessSuite) TestCreateBusiness() {
	tests := []struct {
		in  *typedb.Business
		err error
	}{
		{in: &typedb.Business{
			Username:     "Somewhere",
			Email:        "somewhere@gmail.com",
			Password:     "ExamplePassword",
			BusinessName: "SomewhereBar",
			Type:         "Bar",
			Tags:         nil,
			Description:  "",
			Location: typedb.Location{
				Type:        "Point",
				Coordinates: []float64{100.769652, 13.727892},
			},
			Address:      "Somewhere",
			DisplayImage: "",
			Images:       nil,
			Placement:    nil,
			Menu:         nil,
			Status:       0,
			Verified:     false,
		}, err: nil},
		{in: &typedb.Business{
			Username:     "Somewhere",
			Email:        "somewhere@gmail.com",
			Password:     "ExamplePassword",
			BusinessName: "SomewhereBar",
			Type:         "Bar",
			Tags:         nil,
			Description:  "",
			Location: typedb.Location{
				Type:        "Point",
				Coordinates: []float64{100.769652, 13.727892},
			},
			Address:      "Somewhere",
			DisplayImage: "",
			Images:       nil,
			Placement:    nil,
			Menu:         nil,
			Status:       0,
			Verified:     false,
		}, err: commonErr.INTERNAL},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := b.BusinessDB.CreateBusiness(ctx, tt.in)

		b.Assert().Equal(tt.err, err)
	}

	// check new business
	tmp := &typedb.Business{Username: "Somewhere", Password: "ExamplePassword"}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := b.BusinessDB.AuthenticateBusiness(ctx, tmp)

	b.Assert().Equal(nil, err)
}

func TestBusinessSuite(t *testing.T) {
	suite.Run(t, new(BusinessSuite))
}
