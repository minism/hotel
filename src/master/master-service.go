package master

import (
	"context"
	"log"

	"minornine.com/hotel/src/proto"
)

type MasterService struct{}

func (s *MasterService) RegisterSpawner(context.Context, *hotel_pb.RegisterSpawnerRequest) (*hotel_pb.RegisterSpawnerResponse, error) {
	log.Println("Received RPC from spawner")
	return &hotel_pb.RegisterSpawnerResponse {
		Status: hotel_pb.Status_OK,
	}, nil
}
