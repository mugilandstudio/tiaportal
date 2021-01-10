package grpc

import (
	"log"

	"github.com/mugilandstudio/tiaportal/pkg/server"
	"github.com/spf13/cobra"
)

var (
	srv = server.NewGRPCServer()
	Cmd = &cobra.Command{
		Use:   "grpc",
		Short: "gRPCサーバー関連のコマンドです。",
	}
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "gRPCサーバーを起動します。",
		Run: func(cmd *cobra.Command, args []string) {
			if err := srv.Start(); err != nil {
				log.Fatalf("Failed start gRPC server: %v", err)
			}
		},
	}
)

func init() {
	Cmd.AddCommand(startCmd)
}
