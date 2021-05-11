package placementdb

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"testing"
	"time"
)

const (
	specialTaleID = "5fcde2ec209efa45620a08b6"
)

type PlacementSuite struct {
	suite.Suite
	PlacementDB *PlacementDB
}

func (p *PlacementSuite) SetupSuite() {
	// Create MongoDB client and verify connection
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		p.T().Fatal("Mongo connection URI is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		p.T().Fatal("Could create a mongo client")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		p.T().Fatal("Could not connect to mongodb")
	}

	db := client.Database("test")

	// Attach CustomerDB type to Suite
	p.PlacementDB = &PlacementDB{
		BusCol: db.Collection("business"),
	}
}

func (p *PlacementSuite) TestGetAndUpdate() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// intial get
	initial, err := p.PlacementDB.GetPlacement(ctx, specialTaleID)
	p.Assert().Equal(nil, err)
	p.Assert().Equal(4, len(initial.Seats))

	// update
	err = p.PlacementDB.UpdatePlacement(ctx, specialTaleID, typedb.Placement{
		Width:  800,
		Height: 800,
		Seats:  make([]typedb.Seat, 0),
	})

	p.Assert().Equal(nil, err)

	// get after update
	after, err := p.PlacementDB.GetPlacement(ctx, specialTaleID)
	p.Assert().Equal(nil, err)
	p.Assert().Equal(0, len(after.Seats))
}

func TestBusinessSuite(t *testing.T) {
	suite.Run(t, new(PlacementSuite))
}
