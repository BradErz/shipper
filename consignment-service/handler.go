// consignment-service/handler.go
package main

import (
	"log"

	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"

	pb "github.com/BradErz/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/BradErz/shippy/vessel-service/proto/vessel"
)

type service struct {
	session *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}

	log.Printf("found vessel: %s \n", vesselResponse.Vessel.Name)

	req.VesselId = vesselResponse.Vessel.Id

	err = repo.Create(req)

	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil

}
