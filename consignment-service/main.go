// consignment-service/main.go
package main

import (
	"log"
	"os"

	pb "github.com/BradErz/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/BradErz/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	mongoDbHost = "localhost:27017"
)


func main()  {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = mongoDbHost
	}

	session, err :=  CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("could not connect to the db with host: %s - %v", host, err)
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient :=  vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{vesselClient: vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}