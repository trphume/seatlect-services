package main

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/admindb"
	"github.com/tphume/seatlect-services/internal/database/businessdb"
	"github.com/tphume/seatlect-services/internal/database/requestdb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/admin_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/business_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/request_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/user_api"
	"github.com/tphume/seatlect-services/internal/web/adminwb"
	"github.com/tphume/seatlect-services/internal/web/businesswb"
	"github.com/tphume/seatlect-services/internal/web/requestwb"
	"github.com/tphume/seatlect-services/internal/web/userwb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

func main() {
	// Construct mongo db client and collection
	mongoURI := os.Getenv("MONGO_URI")
	if len(mongoURI) == 0 {
		log.Fatal("missing mongo uri")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("error creating mongo client: ", err.Error())
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("error pinging mongo database")
	}

	adminCol := client.Database("test").Collection("admin")
	busCol := client.Database("test").Collection("business")
	reqCol := client.Database("test").Collection("request")

	// Connect to google cloud and get image bucket
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	imgClient, err := storage.NewClient(ctx, option.WithCredentialsFile("seatlect-image.creds.json"))
	if err != nil {
		log.Fatal("error creating google cloud storage client: ", err.Error())
	}

	imgBucket := imgClient.Bucket("seatlect-images")

	// Construct route handlers and repo
	adminRepo := &admindb.AdminDB{AdminCol: adminCol}
	adminServer := &adminwb.Server{Repo: adminRepo}

	busRepo := &businessdb.BusinessDB{BusCol: busCol, ImageBucket: imgBucket}
	busServer := &businesswb.Server{Repo: busRepo}
	userServer := &userwb.Server{Repo: busRepo}

	reqRepo := &requestdb.RequestDB{ReqCol: reqCol, BusCol: busCol}
	reqServer := &requestwb.Server{Repo: reqRepo}

	// Register routes
	server := echo.New()
	apiV1 := server.Group("/api/v1")

	admin_api.RegisterHandlers(apiV1, adminServer)
	business_api.RegisterHandlers(apiV1, busServer)
	user_api.RegisterHandlers(apiV1, userServer)
	request_api.RegisterHandlers(apiV1, reqServer)

	// Start the server
	log.Println("Starting the server on port 0.0.0.0:9999")
	log.Fatal(server.Start("0.0.0.0:9999"))
}
