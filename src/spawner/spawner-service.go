package spawner

import (
	"context"
	"log"

	"minornine.com/hotel/src/proto"
)

type SpawnerService struct{}

func (s *SpawnerService) CheckStatus(context.Context, *hotel_pb.CheckStatusRequest) (*hotel_pb.CheckStatusResponse, error) {
	log.Println("Received RPC")
	return &hotel_pb.CheckStatusResponse{
		Status: hotel_pb.Status_OK,
	}, nil
}

func (s *SpawnerService) SpawnServer(context.Context, *hotel_pb.SpawnServerRequest) (*hotel_pb.SpawnServerResponse, error) {
	log.Println("Received RPC")
	return &hotel_pb.SpawnServerResponse{
		Status: hotel_pb.Status_OK,
	}, nil
}
