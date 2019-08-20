package master

import (
	"context"
	"fmt"
	"log"

	"minornine.com/hotel/src/proto"

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

	client := proto.NewSpawnerServiceClient(conn)
	response, err := client.Test(context.Background(), &proto.TestRequest{
		Body: "test",
	})
	fmt.Println(response)
}
