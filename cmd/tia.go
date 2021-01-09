package main

import (
	"log"

	"github.com/mugilandstudio/tiaportal/cmd/server"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "tia",
		Short: "tiaportal用コマンドです。",
	}
)

func init() {
	cmd.AddCommand(server.Cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Failed command")
	}
}
