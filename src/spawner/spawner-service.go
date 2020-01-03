package spawner

import (
	"context"
	"log"

	hotel_pb "minornine.com/hotel/src/proto"
)

type SpawnerService struct{}

func (s *SpawnerService) CheckStatus(context.Context, *hotel_pb.CheckStatusRequest) (*hotel_pb.CheckStatusResponse, error) {
	log.Println("Received RPC")
	return &hotel_pb.CheckStatusResponse{}, nil
}

func (s *SpawnerService) SpawnServer(context.Context, *hotel_pb.SpawnServerRequest) (*hotel_pb.SpawnServerResponse, error) {
	log.Println("Received RPC")
	return &hotel_pb.SpawnServerResponse{}, nil
}
