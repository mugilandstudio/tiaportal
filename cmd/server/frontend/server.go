package frontend

import (
	"log"

	"github.com/mugilandstudio/tiaportal/pkg/server"
	"github.com/spf13/cobra"
)

var (
	srv = server.NewFrontendServer()
	Cmd = &cobra.Command{
		Use:   "frontend",
		Short: "frontendサーバー関連のコマンドです。",
	}
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "frontendサーバーを起動します。",
		Run: func(cmd *cobra.Command, args []string) {
			if err := srv.Start(); err != nil {
				log.Fatalf("Failed start frontend server: %v", err)
			}
		},
	}
)

func init() {
	Cmd.AddCommand(startCmd)
}
