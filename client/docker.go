package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

func StartDocker(socketAddr string, disconnectWrite bool) {
	// addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8800")
	// if err != nil {
	// 	panic(err)
	// }
	// c, err := net.DialTCP("tcp", nil, addr)

	addr, err := net.ResolveUnixAddr("unix", socketAddr)
	if err != nil {
		panic(err)
	}
	c, err := net.DialUnix("unix", nil, addr)

	// c, err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(c)
	br := bufio.NewReader(c)

	bw.WriteString("POST /v1.45/containers/nginx/attach?stderr=1&stdout=1&stream=1 HTTP/1.1\r\n")

	bw.WriteString("Host: api.moby.localhost\r\n")
	bw.WriteString("User-Agent: Docker-Client/27.2.1-rd (linux)\r\n")
	bw.WriteString("Content-Length: 0\r\n")
	bw.WriteString("Connection: Upgrade\r\n")
	bw.WriteString("Content-Type: text/plain\r\n")
	bw.WriteString("Upgrade: tcp\r\n")

	bw.WriteString("\r\n")
	if err := bw.Flush(); err != nil {
		panic(err)
	}

	resp, err := http.ReadResponse(br, &http.Request{Method: "POST"})
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 101 {
		fmt.Println(resp.StatusCode)
		panic(websocket.ErrBadStatus)
	}
	if strings.ToLower(resp.Header.Get("Upgrade")) != "tcp" ||
		strings.ToLower(resp.Header.Get("Connection")) != "upgrade" {
		panic(websocket.ErrBadUpgrade)
	}

	if disconnectWrite {
		fmt.Println("Closing write")
		fmt.Println(c.CloseWrite())
	}

	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = br.Read(msg); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received: %s.\n", msg[:n])

		// br.Req.Body.Close()
	}
}
