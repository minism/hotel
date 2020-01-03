package master

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/peer"
	hotel_pb "minornine.com/hotel/src/proto"
)

type MasterService struct{}

func (s *MasterService) RegisterSpawner(ctx context.Context, request *hotel_pb.RegisterSpawnerRequest) (*hotel_pb.RegisterSpawnerResponse, error) {
	pr, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("Received RPC from spawner: %v", pr.Addr)
	}
	host := request.Host
	if host == "" {
		var err error
		host, _, err = net.SplitHostPort(pr.Addr.String())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "No valid host found for request.")
		}
	}

	// Register the spawner with the manager.
	spawner := Spawner{
		Host: host,
		Port: request.Port,
	}
	RegisterSpawner(spawner)

	return &hotel_pb.RegisterSpawnerResponse{}, nil
}
