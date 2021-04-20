package businessdb

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
	"time"
)

type BusinessDB struct {
	BusCol      *mongo.Collection
	ImageBucket *storage.BucketHandle
}

func (b *BusinessDB) ListBusiness(ctx context.Context, searchParams typedb.ListBusinessParams) ([]typedb.Business, error) {
	// Construcut params
	limit := new(int64)
	*limit = int64(searchParams.Limit)

	// TODO: construct sorting option

	// Construct query
	query := bson.D{
		{"location", bson.D{
			{"$geoWithin", bson.M{"$centerSphere": bson.A{searchParams.Location.Coordinates, 0.00156786503}}},
		}},
		{"status", 1},
		{"type", searchParams.Type},
	}

	if searchParams.Name != "" {
		query = bson.D{
			{"$text", bson.M{"$search": searchParams.Name}},
			{"location", bson.D{
				{"$geoWithin", bson.M{"$centerSphere": bson.A{searchParams.Location.Coordinates, 0.00156786503}}},
			}},
			{"status", 1},
			{"type", searchParams.Type},
		}
	}

	// search for business by name
	businesses, err := b.BusCol.Find(
		ctx,
		query,
		options.Find().SetLimit(*limit).SetProjection(bson.M{"placement": 0, "employee": 0}),
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Business
	if err = businesses.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (b *BusinessDB) ListBusinessByIds(ctx context.Context, ids []string) ([]typedb.Business, error) {
	objIds := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objIds[i], _ = primitive.ObjectIDFromHex(id)
	}

	businesses, err := b.BusCol.Find(
		ctx,
		bson.D{
			{"_id",
				bson.D{
					{"$in", objIds},
				},
			},
		},
		options.Find().SetProjection(bson.M{"placement": 0, "employee": 0}),
	)

	if err != nil {
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Business
	if err = businesses.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (b *BusinessDB) AuthenticateBusiness(ctx context.Context, business *typedb.Business) (string, error) {
	res := b.BusCol.FindOne(
		ctx,
		bson.M{"status": 1, "username": business.Username},
		options.FindOne().SetProjection(bson.M{"_id": 1}),
	)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", commonErr.NOTFOUND
		}

		return "", commonErr.INTERNAL
	}

	pw := business.Password
	if err := res.Decode(business); err != nil {
		return "", commonErr.INTERNAL
	}

	if err := bcrypt.CompareHashAndPassword([]byte(business.Password), []byte(pw)); err != nil {
		return "", commonErr.NOTFOUND
	}

	return business.Id.Hex(), nil
}

func (b *BusinessDB) CreateBusiness(ctx context.Context, business *typedb.Business) error {
	pw, err := bcrypt.GenerateFromPassword([]byte(business.Password), 12)
	if err != nil {
		return commonErr.INTERNAL
	}

	business.Password = string(pw)
	business.Id = primitive.NewObjectIDFromTimestamp(time.Now())

	_, err = b.BusCol.InsertOne(ctx, business)
	if err != nil {
		// TODO: better error handling
		return commonErr.INTERNAL
	}

	return nil
}

func (b *BusinessDB) SimpleListBusiness(ctx context.Context, status int, page int) ([]typedb.Business, error) {
	if page == 0 {
		page = 1
	}

	// Construct params
	limit := int64(10)
	p := int64((page - 1) * 10)

	req, err := b.BusCol.Find(ctx, bson.M{"status": status}, &options.FindOptions{
		Limit: &limit,
		Skip:  &p,
	})

	if err != nil {
		return nil, commonErr.INTERNAL
	}

	var res []typedb.Business
	if err = req.All(ctx, &res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res, nil
}

func (b *BusinessDB) GetBusinessById(ctx context.Context, id string, withMenu bool) (*typedb.Business, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, commonErr.INVALID
	}

	// construct projection
	projection := bson.M{"menu": 0, "placement": 0, "employee": 0}
	if withMenu {
		projection = bson.M{"placement": 0, "employee": 0}
	}

	business := b.BusCol.FindOne(
		ctx,
		bson.M{"_id": pId},
		options.FindOne().SetProjection(projection),
	)

	if business.Err() != nil {
		if business.Err() == mongo.ErrNoDocuments {
			return nil, commonErr.NOTFOUND
		}

		return nil, commonErr.INTERNAL
	}

	var res typedb.Business
	if err = business.Decode(&res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return &res, nil
}

func (b *BusinessDB) UpdateBusinessById(ctx context.Context, business typedb.Business) error {
	res, err := b.BusCol.UpdateOne(
		ctx,
		bson.M{"_id": business.Id},
		bson.D{
			{"$set",
				bson.D{
					{"businessName", business.BusinessName},
					{"type", business.Type},
					{"tags", business.Tags},
					{"description", business.Description},
					{"location", business.Location},
					{"address", business.Address},
				}},
		},
	)
	if err != nil {
		return commonErr.INTERNAL
	}

	if res.ModifiedCount == 0 {
		return commonErr.NOTFOUND
	}

	return nil
}

func (b *BusinessDB) UpdateBusinessDIById(ctx context.Context, id string, image string) (string, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", commonErr.INVALID
	}

	// Upload Image to bucket
	wr := b.ImageBucket.Object(uuid.NewString()).NewWriter(ctx)

	index := strings.Index(image, ",")
	imgDecoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image[index+1:]))

	if _, err = io.Copy(wr, imgDecoder); err != nil {
		return "", commonErr.INVALID
	}

	if err = wr.Close(); err != nil {
		return "", commonErr.INVALID
	}

	attrs, err := b.ImageBucket.Attrs(ctx)
	if err != nil {
		return "", commonErr.INTERNAL
	}

	image = fmt.Sprintf("https://storage.googleapis.com/%s/%s", attrs.Name, wr.Name)

	// Update in Mongo
	res := b.BusCol.FindOneAndUpdate(
		ctx,
		bson.M{"_id": pId},
		bson.D{
			{
				"$set",
				bson.D{{"displayImage", image}},
			},
		},
	)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", commonErr.NOTFOUND
		}

		return "", commonErr.INTERNAL
	}

	return image, nil
}

func (b *BusinessDB) AppendBusinessImage(ctx context.Context, id string, image string) (string, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", commonErr.INVALID
	}

	// Upload Image to bucket
	wr := b.ImageBucket.Object(uuid.NewString()).NewWriter(ctx)

	index := strings.Index(image, ",")
	imgDecoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image[index+1:]))

	if _, err = io.Copy(wr, imgDecoder); err != nil {
		return "", commonErr.INVALID
	}

	if err = wr.Close(); err != nil {
		return "", commonErr.INVALID
	}

	attrs, err := b.ImageBucket.Attrs(ctx)
	if err != nil {
		return "", commonErr.INTERNAL
	}

	image = fmt.Sprintf("https://storage.googleapis.com/%s/%s", attrs.Name, wr.Name)

	// Update in Mongo
	res := b.BusCol.FindOneAndUpdate(
		ctx,
		bson.M{"_id": pId},
		bson.D{
			{"$push", bson.D{
				{"images", image},
			}},
		},
	)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", commonErr.NOTFOUND
		}

		return "", commonErr.INTERNAL
	}

	return image, nil
}

func (b *BusinessDB) RemoveBusinessImage(ctx context.Context, id string, pos int) error {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	// Non-atomic version
	if res := b.BusCol.FindOneAndUpdate(ctx, bson.M{"_id": pId}, bson.M{"$unset": bson.M{fmt.Sprintf("images.%d", pos): 1}}); res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	if res := b.BusCol.FindOneAndUpdate(ctx, bson.M{"_id": pId}, bson.M{"$pull": bson.M{"images": nil}}); res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	return nil
}

func (b *BusinessDB) ListMenuItem(ctx context.Context, id string) ([]typedb.MenuItems, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, commonErr.INVALID
	}

	items := b.BusCol.FindOne(
		ctx,
		bson.M{"_id": pId},
		options.FindOne().SetProjection(bson.M{"menu": 1}),
	)

	if items.Err() != nil {
		if items.Err() == mongo.ErrNoDocuments {
			return nil, commonErr.NOTFOUND
		}

		return nil, commonErr.INTERNAL
	}

	var res typedb.Business
	if err = items.Decode(&res); err != nil {
		return nil, commonErr.INTERNAL
	}

	return res.Menu, nil
}

func (b *BusinessDB) AppendMenuItem(ctx context.Context, id string, item typedb.MenuItems) (string, error) {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", commonErr.INVALID
	}

	// Upload image to bucket
	wr := b.ImageBucket.Object(uuid.NewString()).NewWriter(ctx)

	index := strings.Index(item.Image, ",")
	imgDecoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(item.Image[index+1:]))

	if _, err = io.Copy(wr, imgDecoder); err != nil {
		return "", commonErr.INVALID
	}

	if err = wr.Close(); err != nil {
		return "", commonErr.INVALID
	}

	attrs, err := b.ImageBucket.Attrs(ctx)
	if err != nil {
		return "", commonErr.INTERNAL
	}

	item.Image = fmt.Sprintf("https://storage.googleapis.com/%s/%s", attrs.Name, wr.Name)

	// Add to mongo
	res, err := b.BusCol.UpdateOne(
		ctx,
		bson.D{
			{"_id", pId},
			{"menu", bson.D{
				{"$ne", item.Name},
			}},
		},
		bson.D{
			{"$push", bson.D{
				{"menu", item},
			}},
		},
	)

	if err != nil {
		return "", commonErr.INTERNAL
	}

	// Actually can be the case that name is duplicate
	if res.MatchedCount == 0 {
		return "", commonErr.NOTFOUND
	}

	return item.Image, nil
}

func (b *BusinessDB) RemoveMenuItem(ctx context.Context, id string, name string) error {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	res, err := b.BusCol.UpdateOne(
		ctx,
		bson.M{"_id": pId},
		bson.D{
			{"$pull",
				bson.D{
					{"menu",
						bson.D{
							{"name", name},
						},
					},
				},
			},
		},
	)

	if err != nil {
		return commonErr.INTERNAL
	}

	if res.ModifiedCount == 0 {
		return commonErr.NOTFOUND
	}

	return nil
}

func (b *BusinessDB) DeleteBusiness(ctx context.Context, id string) error {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	res := b.BusCol.FindOneAndDelete(ctx, bson.M{"_id": pId})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return commonErr.NOTFOUND
		}

		return commonErr.INTERNAL
	}

	return nil
}

func (b *BusinessDB) UpdateBusinessStatus(ctx context.Context, id string, status int) error {
	pId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return commonErr.INVALID
	}

	res, err := b.BusCol.UpdateOne(
		ctx,
		bson.M{"_id": pId},
		bson.D{
			{"$set",
				bson.D{
					{"status", status},
				},
			},
		},
	)

	if err != nil {
		return commonErr.INTERNAL
	}

	if res.ModifiedCount == 0 {
		return commonErr.NOTFOUND
	}

	return nil
}
