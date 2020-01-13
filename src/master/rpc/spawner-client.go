package rpc

import (
	"context"
	"log"

	"github.com/minism/hotel/src/master/models"
	hotel_pb "github.com/minism/hotel/src/proto"

	"google.golang.org/grpc"
)

// SendSpawnServerRequest makes an RPC to a spawner to spawn a game server instance.
func SendSpawnServerRequest(spawner models.Spawner) (*hotel_pb.SpawnServerResponse, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(spawner.Address(), opts...)

	if err != nil {
		log.Printf("Error connecting to RPC host %v: %v", spawner.Address(), err)
		return nil, err
	}
	defer conn.Close()

	client := hotel_pb.NewSpawnerServiceClient(conn)
	response, err := client.SpawnServer(context.Background(), &hotel_pb.SpawnServerRequest{})
	return response, err
}

// SendCheckStatusRequest asks the given spawner to report its current status.
func SendCheckStatusRequest(spawner models.Spawner) (*hotel_pb.SpawnerStatus, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(spawner.Address(), opts...)

	if err != nil {
		log.Printf("Error connecting to RPC host %v: %v", spawner.Address(), err)
		return nil, err
	}
	defer conn.Close()

	client := hotel_pb.NewSpawnerServiceClient(conn)
	response, err := client.CheckStatus(context.Background(), &hotel_pb.CheckStatusRequest{})
	if err != nil {
		return nil, err
	}
	return response.Status, nil
}
