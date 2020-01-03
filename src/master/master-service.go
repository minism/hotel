package master

import (
	"context"
	"log"

	"google.golang.org/grpc/peer"
	hotel_pb "minornine.com/hotel/src/proto"
)

type MasterService struct{}

func (s *MasterService) RegisterSpawner(ctx context.Context, req *hotel_pb.RegisterSpawnerRequest) (*hotel_pb.RegisterSpawnerResponse, error) {
	log.Println("Received RPC from spawner")
	pr, ok := peer.FromContext(ctx)
	if ok {
		log.Println(pr.Addr.String)
	}
	return &hotel_pb.RegisterSpawnerResponse{
		Status: hotel_pb.Status_OK,
	}, nil
}
