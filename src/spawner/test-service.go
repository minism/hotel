package spawner

import (
	"context"
	"log"

	"minornine.com/hotel/src/proto"
)

type TestService struct{}

func (s *TestService) Test(context.Context, *proto.TestRequest) (*proto.TestResponse, error) {
	log.Println("Received RPC")
	return &proto.TestResponse{
		Status: proto.Status_OK,
	}, nil
}
