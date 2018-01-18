package main

import (
	"log"

	pb "github.com/BradErz/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"

	"context"
	"errors"
)

// Repository - contains availiable methods for vessel-service
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository - contains a list of vessels
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error)  {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("no vessel found by that spec")
}

// gRPC service handler
type service struct {
	repo Repository
}

// FindAvailable
func (s *service) FindAvailiable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main()  {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels: vessels}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo: repo})

	if err := srv.Run() ;err != nil {
		log.Fatalf("error running service: %v", err)
	}
}