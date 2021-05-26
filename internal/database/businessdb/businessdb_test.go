package businessdb

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/api/option"
	"os"
	"testing"
	"time"
)

const (
	brightioID   = "5facafef6b28446f285d7ae4"
	beerBurgerId = "5facaff31c6d49b2c7256bf3"
	ironBuffetId = "5facaff9e4d46967c9c2a558"
	admin1ID     = "604dfa455226a8714411f33d"
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

	// Setup google cloud storage client
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	stClient, err := storage.NewClient(ctx, option.WithCredentialsFile("../../../seatlect-image.creds.json"))
	if err != nil {
		b.T().Fatal("Error creating google cloud storage client: " + err.Error())
	}

	b.BusinessDB.ImageBucket = stClient.Bucket("seatlect-images-test")
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

func (b *BusinessSuite) TestCreateAndDelete() {
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
			Placement:    typedb.Placement{},
			Menu:         nil,
			Status:       1,
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
			Placement:    typedb.Placement{},
			Menu:         nil,
			Status:       1,
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
	id, err := b.BusinessDB.AuthenticateBusiness(ctx, tmp)

	b.Assert().Equal(nil, err)

	// now delete
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = b.BusinessDB.DeleteBusiness(ctx, id)

	b.Assert().Equal(nil, err)
}

func (b *BusinessSuite) TestSimpleListBusiness() {
	tests := []struct {
		status int
		page   int
		lenout int
		err    error
	}{
		{status: 1, page: 1, lenout: 4, err: nil},
		{status: 1, page: -1, lenout: 0, err: commonErr.INTERNAL},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		out, err := b.BusinessDB.SimpleListBusiness(ctx, tt.status, tt.page)

		b.Assert().Equal(tt.err, err)
		b.Assert().Equal(tt.lenout, len(out))
	}
}

func (b *BusinessSuite) TestGetBusinessById() {
	tests := []struct {
		in      string
		nameout string
		err     error
	}{
		{in: brightioID, nameout: "Brightio", err: nil},
		{in: "randomID", nameout: "", err: commonErr.INVALID},
		{in: admin1ID, nameout: "", err: commonErr.NOTFOUND},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		out, err := b.BusinessDB.GetBusinessById(ctx, tt.in, false)

		b.Assert().Equal(tt.err, err)
		if err == nil {
			b.Assert().Equal(tt.nameout, out.BusinessName)
		}
	}
}

func (b *BusinessSuite) TestUpdateBusinessById() {
	pBrightioId, _ := primitive.ObjectIDFromHex(brightioID)
	pAdminId, _ := primitive.ObjectIDFromHex(admin1ID)

	tests := []struct {
		in  typedb.Business
		err error
	}{
		{in: typedb.Business{
			Id:           pBrightioId,
			BusinessName: "Brightio",
			Type:         "Bar",
			Tags:         []string{"Hello"},
			Description:  "Something idk",
			Location: typedb.Location{
				Type:        "Point",
				Coordinates: []float64{100.769652, 13.727892},
			},
			Address: "Somewhere",
		}, err: nil},
		{in: typedb.Business{Id: pAdminId}, err: commonErr.NOTFOUND},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := b.BusinessDB.UpdateBusinessById(ctx, tt.in)

		b.Assert().Equal(tt.err, err)
	}
}

func (b *BusinessSuite) TestUpdateBusinessDIById() {
	tests := []struct {
		in  string
		img string
		err error
	}{
		{in: brightioID, img: b64Img, err: nil},
		{in: admin1ID, img: b64Img, err: commonErr.NOTFOUND},
		{in: admin1ID, img: "wrong", err: commonErr.INVALID},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err := b.BusinessDB.UpdateBusinessDIById(ctx, tt.in, tt.img)

		b.Assert().Equal(tt.err, err)
	}
}

func (b *BusinessSuite) TestListMenuItem() {
	tests := []struct {
		in     string
		lenout int
		err    error
	}{
		{in: beerBurgerId, lenout: 2, err: nil},
		{in: brightioID, lenout: 0, err: nil},
		{in: admin1ID, lenout: 0, err: commonErr.NOTFOUND},
		{in: "randomid", lenout: 0, err: commonErr.INVALID},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		out, err := b.BusinessDB.ListMenuItem(ctx, tt.in)

		b.Assert().Equal(tt.err, err)
		b.Assert().Equal(tt.lenout, len(out))
	}
}

func (b *BusinessSuite) TestAppendMenuItem() {
	tests := []struct {
		in   string
		item typedb.MenuItems
		err  error
	}{
		{in: ironBuffetId, item: typedb.MenuItems{
			Name:        "Pisco",
			Description: "A delicate Pisco cat dish",
			Image:       b64Img,
			Price:       100,
		}, err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err := b.BusinessDB.AppendMenuItem(ctx, tt.in, tt.item)

		b.Assert().Equal(nil, err)
	}
}

func (b *BusinessSuite) TestRemoveMenuItem() {
	tests := []struct {
		id   string
		name string
		err  error
	}{
		{id: ironBuffetId, name: "Fries", err: nil},
		{id: ironBuffetId, name: "Fries", err: commonErr.NOTFOUND},
		{id: admin1ID, name: "Salty Fries", err: commonErr.NOTFOUND},
		{id: "fewfew", name: "", err: commonErr.INVALID},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := b.BusinessDB.RemoveMenuItem(ctx, tt.id, tt.name)

		b.Assert().Equal(tt.err, err)
	}
}

func (b *BusinessSuite) TestUpdateBusinessStatus() {
	tests := []struct {
		in     string
		status int
		err    error
	}{
		{in: "randomfakeid", status: 100, err: commonErr.INVALID},
		{in: admin1ID, status: 100, err: commonErr.NOTFOUND},
		{in: brightioID, status: 100, err: nil},
	}

	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err := b.BusinessDB.UpdateBusinessStatus(ctx, tt.in, tt.status)

		b.Assert().Equal(tt.err, err)
	}

	// check updated status
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	brightio, err := b.BusinessDB.GetBusinessById(ctx, brightioID, false)

	b.Assert().Equal(nil, err)
	b.Assert().Equal(100, brightio.Status)

	// cleanup
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	_, err = b.BusinessDB.UpdateBusinessStatus(ctx, brightioID, 1)

	b.Assert().Equal(nil, err)
}

func TestBusinessSuite(t *testing.T) {
	suite.Run(t, new(BusinessSuite))
}
