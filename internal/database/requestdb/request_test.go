package requestdb

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type RequestSuite struct {
	suite.Suite
	RequestDB *RequestDB
}

func (r *RequestSuite) SetupSuite() {
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
	r.RequestDB = &RequestDB{
		ReqCol: db.Collection("request"),
		BusCol: db.Collection("business"),
	}
}

func (r *RequestSuite) TestListRequest() {
	tests := []struct {
		in     int
		lenout int
		idout  string
		err    error
	}{
		{in: 0, lenout: 1, idout: brightioID, err: nil},
		{in: 1, lenout: 1, idout: brightioID, err: nil},
		{in: 2, lenout: 0, idout: "", err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		out, err := r.RequestDB.ListRequest(ctx, tt.in)

		r.Assert().Equal(tt.err, err)
		r.Assert().Equal(tt.lenout, len(out))

		if len(out) != 0 {
			r.Assert().Equal(tt.idout, out[0].Id.Hex())
		}
	}
}

func (r *RequestSuite) TestApproveRequest() {
	pBrightioId, _ := primitive.ObjectIDFromHex(brightioID)

	tests := []struct {
		in  string
		err error
	}{
		{in: brightioID, err: nil},
		{in: brightioID, err: commonErr.NOTFOUND},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		_, err := r.RequestDB.ApproveRequest(ctx, tt.in)

		r.Assert().Equal(tt.err, err)
	}

	// recreate apprpove request
	req := &typedb.Request{
		Id:           pBrightioId,
		BusinessName: "Brightio",
		Type:         "Cool Bar",
		Tags:         []string{"BAR", "JAPANESE", "LIVE MUSIC"},
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		Location: typedb.Location{
			Type:        "Point",
			Coordinates: []float64{100.769652, 13.727892},
		},
		Address:   "Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520",
		CreatedAt: time.Now(),
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := r.RequestDB.CreateRequest(ctx, req); err != nil {
		r.FailNow("error on recreate request")
	}
}

func (r *RequestSuite) TestGetRequestById() {
	pBrightioId, _ := primitive.ObjectIDFromHex(brightioID)
	pBeerBurgerId, _ := primitive.ObjectIDFromHex(beerBurgerId)

	tests := []struct {
		in  *typedb.Request
		err error
	}{
		{in: &typedb.Request{Id: pBrightioId}, err: nil},
		{in: &typedb.Request{Id: pBeerBurgerId}, err: commonErr.NOTFOUND},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := r.RequestDB.GetRequestById(ctx, tt.in)

		r.Assert().Equal(tt.err, err)
	}
}

func (r *RequestSuite) TestCreateAndDeleteRequest() {
	pBeerBurgerId, _ := primitive.ObjectIDFromHex(beerBurgerId)

	tests := []struct {
		in  *typedb.Request
		err error
	}{
		{in: &typedb.Request{
			Id:           pBeerBurgerId,
			BusinessName: "BeerBureger",
			Type:         "Super Cool Bar",
			Tags:         []string{},
			Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			Location: typedb.Location{
				Type:        "Point",
				Coordinates: []float64{100.769652, 13.727892},
			},
			Address:   "Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520",
			CreatedAt: time.Now(),
		}, err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		err := r.RequestDB.CreateRequest(ctx, tt.in)

		r.Assert().Equal(tt.err, err)

		// Check created request
		tmp := &typedb.Request{Id: pBeerBurgerId}
		err = r.RequestDB.GetRequestById(ctx, tmp)

		r.Assert().Equal(tt.in.Type, tmp.Type)

		// Delete request
		err = r.RequestDB.DeleteRequest(ctx, beerBurgerId)
		r.Assert().Equal(err, nil)

		err = r.RequestDB.DeleteRequest(ctx, beerBurgerId)
		r.Assert().Equal(err, commonErr.NOTFOUND)
	}
}

func TestRequestSuite(t *testing.T) {
	suite.Run(t, new(RequestSuite))
}
