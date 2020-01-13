package spawner

import (
	"context"
	"log"

	hotel_pb "github.com/minism/hotel/src/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// MasterClient manages the RPC connection to the master server.
type MasterClient struct {
	Addr string
}

// NewMasterClient initializes and returns a MasterClient.
func NewMasterClient(addr string) MasterClient {
	return MasterClient{
		Addr: addr,
	}
}

// Register makes an RPC to the master server to register this spawner instance.
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
	}

	log.Println("Registered successfully with master service.")
	return nil
}
