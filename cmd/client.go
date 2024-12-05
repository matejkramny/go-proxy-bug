package cmd

import (
	"fmt"

	"github.com/matejkramny/go-proxy-bug/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "connect to server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
		client.StartRaw()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
