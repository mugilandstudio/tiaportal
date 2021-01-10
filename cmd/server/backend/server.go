package backend

import (
	"log"

	"github.com/mugilandstudio/tiaportal/pkg/server"
	"github.com/spf13/cobra"
)

var (
	srv = server.NewBackendServer()
	Cmd = &cobra.Command{
		Use:   "backend",
		Short: "backendサーバー関連のコマンドです。",
	}
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "backendサーバーを起動します。",
		Run: func(cmd *cobra.Command, args []string) {
			if err := srv.Start(); err != nil {
				log.Fatalf("Failed start backend server: %v", err)
			}
		},
	}
)

func init() {
	Cmd.AddCommand(startCmd)
}
