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

type Server struct {
}

// New はコンストラクタです。
func New() *Server {
	return &Server{}
}

// Start はServerを開始します。
func (s *Server) Start() error {
	log.Printf("Start Server")
	srv := grpc.NewServer()
	if err := s.initializeServices(srv); err != nil {
		return err
	}
	ln, err := net.Listen("tcp", ":8081")
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

func (s *Server) initializeServices(srv *grpc.Server) error {
	if err := service.RegisterSampleService(srv); err != nil {
		return err
	}
	return nil
}
