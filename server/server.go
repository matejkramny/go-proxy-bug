package server

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	fmt.Println("rcv")
	ws.Request().Body.Close()
	time.Sleep(time.Second)

	str := "server closed request body.."
	fmt.Printf("Sending: %v\n", str)
	ws.Write([]byte(str))
	time.Sleep(time.Second * 5)

	str = "this isn't received?"
	fmt.Printf("Sending: %v\n", str)
	ws.Write([]byte(str))
	time.Sleep(time.Second * 5)

	// io.Copy(ws, ws)

	fmt.Println("done")
}

// This example demonstrates a trivial echo server.
func Start() {
	http.Handle("/", websocket.Handler(EchoServer))
	fmt.Println("Listening on :8899")
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
