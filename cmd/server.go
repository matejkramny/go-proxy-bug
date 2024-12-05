package cmd

import (
	"fmt"

	"github.com/matejkramny/go-proxy-bug/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts websocket server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
