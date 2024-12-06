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
}

var rawDisconnectWrite bool
var connectDirectly bool
var clientRawCmd = &cobra.Command{
	Use:   "raw",
	Short: "send raw http to the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
		client.StartRaw(rawDisconnectWrite, connectDirectly)
	},
}

var dockerSocket string
var disconnectWrite bool
var clientDockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "connect to docker via unix socket",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client docker called")
		client.StartDocker(dockerSocket, disconnectWrite)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(clientDockerCmd)
	clientCmd.AddCommand(clientRawCmd)

	clientRawCmd.Flags().BoolVarP(&rawDisconnectWrite, "disconnect", "d", false, "disconnect output after connecting to the websocket")
	clientRawCmd.Flags().BoolVarP(&connectDirectly, "direct", "s", false, "connect to the websocket server directly instead of going to proxy")

	clientDockerCmd.Flags().StringVarP(&dockerSocket, "socket", "s", "/var/run/docker.sock", "location of unix socket")
	clientDockerCmd.Flags().BoolVarP(&disconnectWrite, "disconnect", "d", false, "disconnect output after connecting to the websocket")
}
