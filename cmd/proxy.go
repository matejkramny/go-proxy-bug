package cmd

import (
	"fmt"

	"github.com/matejkramny/go-proxy-bug/proxy"
	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "starts a proxy server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proxy called")
		proxy.Serve()
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
