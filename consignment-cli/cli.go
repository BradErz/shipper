// consignment-cli/cli.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/BradErz/shippy/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main()  {
	// Setup a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)

	// Contact the server and print the response
	consignment, err := parseFile(defaultFilename)

	if err != nil {
		log.Fatalf("failed to parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("failed to create consignment: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err :=  client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to get consignments: %v", err)
	}

	for key, value := range getAll.Consignments {
		fmt.Printf("%d: %v", key, value)
	}
}