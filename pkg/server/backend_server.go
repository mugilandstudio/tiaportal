package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/mugilandstudio/tiaportal/proto"
	"google.golang.org/grpc"
)

type BackendServer struct{}

func NewBackendServer() *BackendServer {
	return &BackendServer{}
}

func (s *BackendServer) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterSampleServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts); err != nil {
		return err
	}
	return http.ListenAndServe(":8081", mux)
}
