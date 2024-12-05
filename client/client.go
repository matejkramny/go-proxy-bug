package client

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func Start() {
	origin := "http://localhost/"
	url := "ws://localhost:8899/"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	// if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
	// 	log.Fatal(err)
	// }

	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received: %s.\n", msg[:n])
		ws.Req.Body.Close()
	}
}
