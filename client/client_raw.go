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

func StartRaw() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8800")
	if err != nil {
		panic(err)
	}
	c, err := net.DialTCP("tcp", nil, addr)
	// c, err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(c)
	br := bufio.NewReader(c)

	bw.WriteString("GET / HTTP/1.1\r\n")

	// According to RFC 6874, an HTTP client, proxy, or other
	// intermediary must remove any IPv6 zone identifier attached
	// to an outgoing URI.
	bw.WriteString("Host: localost:8800\r\n")
	bw.WriteString("Upgrade: websocket\r\n")
	bw.WriteString("Connection: Upgrade\r\n")
	bw.WriteString("Sec-WebSocket-Key: nonceonce\r\n")
	bw.WriteString("Origin: http://localhost/\r\n")
	bw.WriteString("Sec-WebSocket-Version: 13\r\n")
	// if len(config.Protocol) > 0 {
	// 	bw.WriteString("Sec-WebSocket-Protocol: " + strings.Join(config.Protocol, ", ") + "\r\n")
	// }
	// TODO(ukai): send Sec-WebSocket-Extensions.
	// err = config.Header.WriteSubset(bw, handshakeHeader)
	// if err != nil {
	// 	return err
	// }

	bw.WriteString("\r\n")
	if err := bw.Flush(); err != nil {
		panic(err)
	}

	resp, err := http.ReadResponse(br, &http.Request{Method: "GET"})
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 101 {
		fmt.Println(resp.StatusCode)
		panic(websocket.ErrBadStatus)
	}
	if strings.ToLower(resp.Header.Get("Upgrade")) != "websocket" ||
		strings.ToLower(resp.Header.Get("Connection")) != "upgrade" {
		panic(websocket.ErrBadUpgrade)
	}

	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = br.Read(msg); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received: %s.\n", msg[:n])
		fmt.Println("Closing write")
		fmt.Println(c.CloseWrite())

		// br.Req.Body.Close()
	}
}
