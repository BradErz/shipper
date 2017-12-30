package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/BradErz/shipper/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	address  = "localhost:50051"
	fileName = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	cmd.Init()

	// Create new greeter client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)
	consignment, err := parseFile(fileName)
	if err != nil {
		log.Fatalf("couldn't parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("couldn't create consignment: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("couldn't list consignments: %v", err)
	}
	for _, value := range getAll.Consignments {
		log.Println(value)
	}
}
