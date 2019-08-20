package master

import (
	"context"
	"fmt"
	"log"

	"minornine.com/hotel/src/proto"

	"google.golang.org/grpc"
)

func SendTestSpawnerRequest() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:3001", opts...)

	if err != nil {
		log.Printf("Error connecting to RPC host: %v", err)
		return
	}
	defer conn.Close()

	client := proto.NewSpawnerServiceClient(conn)
	response, err := client.Test(context.Background(), &proto.TestRequest{
		Body: "test",
	})
	fmt.Println(response)
}
