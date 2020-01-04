package spawner

import (
	"context"
	"log"

	hotel_pb "minornine.com/hotel/src/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type MasterClient struct {
	Addr string
}

func NewMasterClient(addr string) MasterClient {
	return MasterClient{
		Addr: addr,
	}
}

func (c *MasterClient) Register(port uint32, spawnerStatus hotel_pb.SpawnerStatus) error {
	// TODO: Credentials.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	// TODO: Consider keeping this connection open with the object lifetime.
	conn, err := grpc.Dial(c.Addr, opts...)
	if err != nil {
		log.Printf("Error connecting to RPC host %v: %v", c.Addr, err)
		return err
	}
	defer conn.Close()

	client := hotel_pb.NewMasterServiceClient(conn)
	request := hotel_pb.RegisterSpawnerRequest{
		Port:   port,
		Status: &spawnerStatus,
	}
	_, err = client.RegisterSpawner(context.Background(), &request)
	st := status.Convert(err)
	if st.Err() != nil {
		log.Printf("Error making master RPC: %v", st.Err())
		return err
	} else {
		log.Printf("OK response from master service")
		return nil
	}
}
