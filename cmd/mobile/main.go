package main

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/businessdb"
	"github.com/tphume/seatlect-services/internal/database/customerdb"
	"github.com/tphume/seatlect-services/internal/database/orderdb"
	"github.com/tphume/seatlect-services/internal/database/reservationdb"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/businesspb"
	"github.com/tphume/seatlect-services/internal/genproto/orderpb"
	"github.com/tphume/seatlect-services/internal/genproto/reservationpb"
	"github.com/tphume/seatlect-services/internal/genproto/userpb"
	"github.com/tphume/seatlect-services/internal/mobile/businessmb"
	"github.com/tphume/seatlect-services/internal/mobile/ordermb"
	"github.com/tphume/seatlect-services/internal/mobile/reservationmb"
	"github.com/tphume/seatlect-services/internal/mobile/usermb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// Construct mongo db client and collection
	mongoURI := os.Getenv("MONGO_URI")
	if len(mongoURI) == 0 {
		log.Fatal("missing mongo uri")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("error creating mongo client: ", err.Error())
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("error pinging mongo database")
	}

	cusCol := client.Database("test").Collection("customer")
	busCol := client.Database("test").Collection("business")
	ordCol := client.Database("test").Collection("order")
	resCol := client.Database("test").Collection("reservation")

	// Construct each mobile api route
	cusRepo := &customerdb.CustomerDB{
		CusCol: cusCol,
		BusCol: busCol,
	}
	userServer := &usermb.Server{Repo: cusRepo}

	busRepo := &businessdb.BusinessDB{BusCol: busCol}
	busServer := &businessmb.Server{Repo: busRepo}

	ordRepo := &orderdb.OrderDB{OrdCol: ordCol}
	ordServer := &ordermb.Server{Repo: ordRepo}

	resRepo := &reservationdb.ReservationDB{ResCol: resCol, BusCol: busCol, OrdCol: ordCol}
	resServer := &reservationmb.Server{Repo: resRepo, SubscribersChannel: make(map[string]map[chan typedb.ReservationPlacement]bool)}

	// Setup the gRPC server
	lis, err := net.Listen("tcp", "0.0.0.0:9700")
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userServer)
	businesspb.RegisterBusinessServiceServer(grpcServer, busServer)
	orderpb.RegisterOrderServiceServer(grpcServer, ordServer)
	reservationpb.RegisterReservationServiceServer(grpcServer, resServer)

	log.Println("Starting server on 0.0.0.0:9700")
	log.Fatal(grpcServer.Serve(lis))
}
