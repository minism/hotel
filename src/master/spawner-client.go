package master

import (
	"context"
	"fmt"
	"log"

	hotel_pb "minornine.com/hotel/src/proto"

	"google.golang.org/grpc"
)

const (
	// Use the docker service name.
	SPAWNER_ADDRESS = "spawner:3001"
)

func SendSpawnServerRequest(spawner *Spawner) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(spawner.Address(), opts...)

	if err != nil {
		log.Printf("Error connecting to RPC host %v: %v", spawner.Address(), err)
		return err
	}
	defer conn.Close()

	client := hotel_pb.NewSpawnerServiceClient(conn)
	// response, err := client.CheckStatus(context.Background(), &hotel_pb.CheckStatusRequest{})
	response, err := client.SpawnServer(context.Background(), &hotel_pb.SpawnServerRequest{})
	fmt.Println(response)
	return err
}

func SendCheckStatusRequest(spawner *Spawner) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(spawner.Address(), opts...)

	if err != nil {
		log.Printf("Error connecting to RPC host %v: %v", spawner.Address(), err)
		return err
	}
	defer conn.Close()

	client := hotel_pb.NewSpawnerServiceClient(conn)
	_, err = client.CheckStatus(context.Background(), &hotel_pb.CheckStatusRequest{})
	return err
}
