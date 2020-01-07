package master

import (
	"context"
	"log"
	"net"

	"minornine.com/hotel/src/shared"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/peer"
	hotel_pb "minornine.com/hotel/src/proto"
)

// MasterService contains RPC implementations for the master service.
type MasterService struct{}

// RegisterSpawner adds the requesting spawner to the database and makes
// it available to clients for requesting spawns.
func (s *MasterService) RegisterSpawner(ctx context.Context, request *hotel_pb.RegisterSpawnerRequest) (*hotel_pb.RegisterSpawnerResponse, error) {
	pr, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("Received RPC from spawner: %v.\n  GameID: %v\n  Max servers: %v", pr.Addr, request.Status.SupportedGameId, request.Status.MaxGameServers)
		log.Printf("")
	}
	host := request.Host
	if host == "" {
		var err error
		host, _, err = net.SplitHostPort(pr.Addr.String())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "No valid host found for request.")
		}
		if len(host) <= 3 {
			host = "localhost"
		}
	}

	// Register the spawner with the manager.
	spawner := Spawner{
		Host:           host,
		Port:           request.Port,
		GameID:         shared.GameIDType(request.Status.SupportedGameId),
		NumGameServers: request.Status.NumGameServers,
		MaxGameServers: request.Status.MaxGameServers,
	}
	RegisterSpawner(spawner)

	return &hotel_pb.RegisterSpawnerResponse{}, nil
}
