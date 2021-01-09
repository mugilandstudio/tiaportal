package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:
		log.Printf("Finish Server")
		return nil
	}
}
