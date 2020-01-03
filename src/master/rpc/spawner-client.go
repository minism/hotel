package rpc

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

func SendTestSpawnerRequest() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(SPAWNER_ADDRESS, opts...)

	if err != nil {
		log.Printf("Error connecting to RPC host %v: %v", SPAWNER_ADDRESS, err)
		return
	}
	defer conn.Close()

	client := hotel_pb.NewSpawnerServiceClient(conn)
	response, err := client.CheckStatus(context.Background(), &hotel_pb.CheckStatusRequest{})
	fmt.Println(response)
}
