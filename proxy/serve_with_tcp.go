package proxy

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"time"
)

// dial the websocket server
func dialer() (net.Conn, error) {
	return net.Dial("tcp", "127.0.0.1:8899")
}

func ServeWithTCP() error {
	// create inmemory listener
	inmem := NewInmemSocket("inmem", 10)

	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		panic(err)
	}

	termch := make(chan os.Signal, 1)
	signal.Notify(termch, os.Interrupt)
	go func() {
		<-termch
		signal.Stop(termch)

		err := listener.Close()
		if err != nil {
			panic(err)
		}

		err = inmem.Close()
		if err != nil {
			panic(err)
		}
	}()

	// start tcp proxy
	go func() {
		connID := 0
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Printf("Failed to accept connection '%s'\n", err)
				break
			}

			p := NewTCPProxy(conn, dialer, inmem)

			// p.Matcher = matcher
			// p.Replacer = replacer
			connID++

			p.OutputHex = false
			p.Log = ColorLogger{
				Verbose:     true,
				VeryVerbose: true,
				Prefix:      fmt.Sprintf("Connection #%v: ", connID),
				// Color:       *colors,
			}

			go p.Start()
		}
	}()

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "proxy.invalid"

			if req.Body != nil {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(body))
				req.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		},
		Transport: &http.Transport{
			Dial: func(string, string) (net.Conn, error) {
				return dialer()
			},
			DisableCompression: true, // for debugging
		},
		ModifyResponse: func(resp *http.Response) error {
			return nil
		},
	}

	server := &http.Server{
		ReadHeaderTimeout: time.Minute,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			proxy.ServeHTTP(w, req)
		}),
	}

	fmt.Println("Listening")

	err = server.Serve(inmem)
	// err = server.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
