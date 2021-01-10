package server

import (
	"github.com/spf13/cobra"

	"github.com/mugilandstudio/tiaportal/cmd/server/backend"
	"github.com/mugilandstudio/tiaportal/cmd/server/frontend"
	"github.com/mugilandstudio/tiaportal/cmd/server/grpc"
)

var (
	Cmd = &cobra.Command{
		Use:   "server",
		Short: "サーバー関連のコマンドです。",
	}
)

func init() {
	Cmd.AddCommand(grpc.Cmd)
	Cmd.AddCommand(backend.Cmd)
	Cmd.AddCommand(frontend.Cmd)
}
