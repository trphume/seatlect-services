package main

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tphume/seatlect-services/internal/database/admindb"
	"github.com/tphume/seatlect-services/internal/database/businessdb"
	"github.com/tphume/seatlect-services/internal/database/orderdb"
	"github.com/tphume/seatlect-services/internal/database/placementdb"
	"github.com/tphume/seatlect-services/internal/database/requestdb"
	"github.com/tphume/seatlect-services/internal/database/reservationdb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/admin_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/business_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/employee_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/order_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/placement_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/request_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/reservation_api"
	"github.com/tphume/seatlect-services/internal/gen_openapi/user_api"
	"github.com/tphume/seatlect-services/internal/web/adminwb"
	"github.com/tphume/seatlect-services/internal/web/businesswb"
	"github.com/tphume/seatlect-services/internal/web/employeewb"
	"github.com/tphume/seatlect-services/internal/web/orderwb"
	"github.com/tphume/seatlect-services/internal/web/placementwb"
	"github.com/tphume/seatlect-services/internal/web/requestwb"
	"github.com/tphume/seatlect-services/internal/web/reservationwb"
	"github.com/tphume/seatlect-services/internal/web/userwb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/api/option"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"time"
)

func main() {
	// Check env
	mongoURI := os.Getenv("MONGO_URI")
	if len(mongoURI) == 0 {
		log.Fatal("missing mongo uri")
	}

	mailUsername := os.Getenv("MAIL_USERNAME")
	if len(mailUsername) == 0 {
		log.Fatal("missing mongo uri")
	}

	mailPassword := os.Getenv("MAIL_PASSWORD")
	if len(mailPassword) == 0 {
		log.Fatal("missing mongo uri")
	}

	// Construct mongo db client and collection
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
	resCol := client.Database("test").Collection("reservation")
	ordCol := client.Database("test").Collection("order")

	// Connect to google cloud and get image bucket
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	imgClient, err := storage.NewClient(ctx, option.WithCredentialsFile("seatlect-image.creds.json"))
	if err != nil {
		log.Fatal("error creating google cloud storage client: ", err.Error())
	}

	imgBucket := imgClient.Bucket("seatlect-images")

	// Construct mail client
	mailClient := gomail.NewDialer("smtp.gmail.com", 587, mailUsername, mailPassword)

	if closer, err := mailClient.Dial(); err != nil {
		log.Fatal("error dialing gmail smtp server: ", err.Error())
	} else {
		_ = closer.Close()
	}

	// Construct route handlers and repo
	adminRepo := &admindb.AdminDB{AdminCol: adminCol}
	adminServer := &adminwb.Server{Repo: adminRepo}

	busRepo := &businessdb.BusinessDB{BusCol: busCol, ImageBucket: imgBucket}
	busServer := &businesswb.Server{Repo: busRepo}
	userServer := &userwb.Server{Repo: busRepo, Mail: mailClient}
	empServer := &employeewb.Server{Repo: busRepo}

	reqRepo := &requestdb.RequestDB{ReqCol: reqCol, BusCol: busCol}
	reqServer := &requestwb.Server{Repo: reqRepo, BusRepo: busRepo}

	pmtRepo := &placementdb.PlacementDB{BusCol: busCol}
	pmtServer := &placementwb.Server{Repo: pmtRepo}

	resRepo := &reservationdb.ReservationDB{ResCol: resCol, BusCol: busCol}
	resServer := &reservationwb.Server{Repo: resRepo}

	ordRepo := &orderdb.OrderDB{OrdCol: ordCol}
	ordServer := &orderwb.Server{Repo: ordRepo}

	// Register routes
	server := echo.New()
	server.Use(middleware.CORS())
	apiV1 := server.Group("/api/v1")

	admin_api.RegisterHandlers(apiV1, adminServer)
	business_api.RegisterHandlers(apiV1, busServer)
	user_api.RegisterHandlers(apiV1, userServer)
	request_api.RegisterHandlers(apiV1, reqServer)
	placement_api.RegisterHandlers(apiV1, pmtServer)
	reservation_api.RegisterHandlers(apiV1, resServer)
	employee_api.RegisterHandlers(apiV1, empServer)
	order_api.RegisterHandlers(apiV1, ordServer)

	// Start the server
	log.Println("Starting the server on port 0.0.0.0:9999")
	log.Fatal(server.Start("0.0.0.0:9999"))
}
