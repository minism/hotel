package spawner

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hotel_pb "github.com/minism/hotel/src/proto"
)

// SpawnerService contains RPC implementations for the spawner service.
type SpawnerService struct {
	config     *Config
	controller *ServerController
}

// NewSpawnerService creates and initializes a SpawnerService.
func NewSpawnerService(config *Config, controller *ServerController) SpawnerService {
	return SpawnerService{
		config:     config,
		controller: controller,
	}
}

// GetStatus returns the status proto representing the current status of the spawner.
func (s *SpawnerService) GetStatus() hotel_pb.SpawnerStatus {
	return hotel_pb.SpawnerStatus{
		SupportedGameId: string(s.config.SupportedGameID),
		NumGameServers:  uint32(s.controller.NumRunningServers()),
		MaxGameServers:  s.config.MaxGameServers,
	}
}

// CheckStatus returns the current status via RPC.
func (s *SpawnerService) CheckStatus(context.Context, *hotel_pb.CheckStatusRequest) (*hotel_pb.CheckStatusResponse, error) {
	status := s.GetStatus()
	return &hotel_pb.CheckStatusResponse{
		Status: &status,
	}, nil
}

// SpawnServer attempts to spawn a game server via the incoming RPC request.
func (s *SpawnerService) SpawnServer(context.Context, *hotel_pb.SpawnServerRequest) (*hotel_pb.SpawnServerResponse, error) {
	if s.controller.Capacity() < 1 {
		return nil, status.Error(codes.ResourceExhausted, "Spawner is already running its maximum server capacity.")
	}
	port, err := s.controller.SpawnServer()
	if err != nil {
		log.Printf("Error spawning server: %v", err)
		return nil, status.Error(codes.Internal, "Error spawning server.")
	}
	status := s.GetStatus()
	return &hotel_pb.SpawnServerResponse{
		Host:		s.config.Host,
		Port:   port,
		Status: &status,
	}, nil
}
