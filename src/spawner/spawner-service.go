package spawner

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hotel_pb "minornine.com/hotel/src/proto"
)

type SpawnerService struct {
	config     *Config
	controller *ServerController
}

func NewSpawnerService(config *Config, controller *ServerController) SpawnerService {
	return SpawnerService{
		config:     config,
		controller: controller,
	}
}

func (s *SpawnerService) GetStatus() hotel_pb.SpawnerStatus {
	return hotel_pb.SpawnerStatus{
		SupportedGameId: string(s.config.SupportedGameID),
		NumGameServers:  uint32(s.controller.NumRunningServers()),
		MaxGameServers:  s.config.MaxGameServers,
	}
}

func (s *SpawnerService) CheckStatus(context.Context, *hotel_pb.CheckStatusRequest) (*hotel_pb.CheckStatusResponse, error) {
	status := s.GetStatus()
	return &hotel_pb.CheckStatusResponse{
		Status: &status,
	}, nil
}

func (s *SpawnerService) SpawnServer(context.Context, *hotel_pb.SpawnServerRequest) (*hotel_pb.SpawnServerResponse, error) {
	port, err := s.controller.SpawnServer()
	if err != nil {
		log.Printf("Error spawning server: %v", err)
		return nil, status.Error(codes.Internal, "Error spawning server.")
	}
	return &hotel_pb.SpawnServerResponse{
		Port: port,
	}, nil
}
