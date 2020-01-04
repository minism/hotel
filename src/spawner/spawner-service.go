package spawner

import (
	"context"
	"log"

	hotel_pb "minornine.com/hotel/src/proto"
)

type SpawnerService struct {
	config *Config
}

func NewSpawnerService(config *Config) SpawnerService {
	return SpawnerService{
		config: config,
	}
}

func (s *SpawnerService) GetStatus() hotel_pb.SpawnerStatus {
	return hotel_pb.SpawnerStatus{
		SupportedGameId: string(s.config.SupportedGameID),
		// TODO: Implement num game servers.
		NumGameServers: 0,
		MaxGameServers: s.config.MaxGameServers,
	}
}

func (s *SpawnerService) CheckStatus(context.Context, *hotel_pb.CheckStatusRequest) (*hotel_pb.CheckStatusResponse, error) {
	status := s.GetStatus()
	return &hotel_pb.CheckStatusResponse{
		Status: &status,
	}, nil
}

func (s *SpawnerService) SpawnServer(context.Context, *hotel_pb.SpawnServerRequest) (*hotel_pb.SpawnServerResponse, error) {
	log.Println("Received RPC")
	return &hotel_pb.SpawnServerResponse{}, nil
}
