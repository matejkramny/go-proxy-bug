package proxy

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"time"
)

func ServeOriginal() error {
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
	}()

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "proxy.invalid"
		},
		Transport: &http.Transport{
			Dial: func(string, string) (net.Conn, error) {
				return net.Dial("tcp", "127.0.0.1:8899")
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
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
