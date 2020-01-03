package spawner

import (
	"context"
	"log"

	hotel_pb "minornine.com/hotel/src/proto"

	"google.golang.org/grpc"
)

type MasterClient struct {
	Addr string
}

func NewMasterClient(addr string) MasterClient {
	return MasterClient{
		Addr: addr,
	}
}

func (c *MasterClient) Test() {
	// TODO: Credentials.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	// TODO: Consider keeping this connection open with the object lifetime.
	conn, err := grpc.Dial(c.Addr, opts...)
	if err != nil {
		log.Printf("Error connecting to RPC host %v: %v", c.Addr, err)
		return
	}
	defer conn.Close()

	client := hotel_pb.NewMasterServiceClient(conn)
	request := hotel_pb.RegisterSpawnerRequest{}
	response, err := client.RegisterSpawner(context.Background(), &request)
	if err != nil {
		log.Printf("Error making master RPC: %v", err)
	} else {
		log.Printf("Response from master service: %v", response.Status)
	}
}
