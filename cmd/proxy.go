package cmd

import (
	"fmt"

	"github.com/matejkramny/go-proxy-bug/proxy"
	"github.com/spf13/cobra"
)

var tcpProxy bool
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "starts a proxy server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proxy called")
		if tcpProxy {
			proxy.ServeWithTCP()
		} else {
			proxy.ServeOriginal()
		}
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)

	proxyCmd.Flags().BoolVarP(&tcpProxy, "tcp", "t", false, "wrap httputil.ReverseProxy in tcp proxy")
}
