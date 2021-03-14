package requestdb

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

const brightioID = "5facafef6b28446f285d7ae4"

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

func TestRequestSuite(t *testing.T) {
	suite.Run(t, new(RequestSuite))
}
