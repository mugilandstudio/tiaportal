package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mugilandstudio/tiaportal/pkg/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

// Start はServerを開始します。
func (s *GRPCServer) Start() error {
	log.Printf("Start gRPC Server")
	srv := grpc.NewServer()
	if err := s.initializeServices(srv); err != nil {
		return err
	}
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}
	if err := srv.Serve(ln); err != nil {
		return err
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:
		log.Printf("Finish Server")
		return nil
	}
}

func (s *GRPCServer) initializeServices(srv *grpc.Server) error {
	if err := service.RegisterSampleService(srv); err != nil {
		return err
	}
	return nil
}
