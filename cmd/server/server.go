package server

import (
	"github.com/mugilandstudio/tiaportal/pkg/server"
	"github.com/spf13/cobra"
)

var (
	srv = server.New()
	Cmd = &cobra.Command{
		Use:   "server",
		Short: "server関連のコマンドです。",
	}
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "serverを起動します。",
		Run: func(cmd *cobra.Command, args []string) {
			srv.Start()
		},
	}
)

func init() {
	Cmd.AddCommand(startCmd)
}
