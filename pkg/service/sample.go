package service

import (
	"context"

	pb "github.com/mugilandstudio/tiaportal/proto"
	"google.golang.org/grpc"
)

func RegisterSampleService(srv *grpc.Server) error {
	s, err := NewSampleService()
	if err != nil {
		return err
	}
	pb.RegisterSampleServiceServer(srv, s)
	return nil
}

type SampleService struct {
	pb.UnimplementedSampleServiceServer
}

func NewSampleService() (*SampleService, error) {
	return &SampleService{}, nil
}

func (s *SampleService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{}, nil
}
