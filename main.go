package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/armon/go-socks5"
	"gopkg.in/elazarl/goproxy.v1"
)

func main() {
	socksPort := flag.Int("sport", 2222, "The port on which the SOCKS proxy listens for connections")
	socksHost := flag.String("shost", "127.0.0.1", "The host/ip on which the SOCKS proxy listens for connections")
	httpPort := flag.Int("hport", 2223, "The port on which the HTTP proxy listens for connections")
	httpHost := flag.String("hhost", "127.0.0.1", "The host/ip on which the HTTP proxy listens for connections")
	verbose := flag.Bool("verbose", false, "If verbose logging should be done")

	flag.Parse()

	go startSocksProxy(makeListenAddress(*socksHost, *socksPort), *verbose)

	startHTTPProxy(makeListenAddress(*httpHost, *httpPort), *verbose)
}

func startSocksProxy(addr string, verbose bool) {
	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal("Failed to create socks proxy configuration: ", err)
	}

	// Create SOCKS5 proxy on localhost port 2222
	err = server.ListenAndServe("tcp", addr)
	if err != nil {
		log.Fatal("Failed to start socks proxy: ", err)
	}
}

func startHTTPProxy(addr string, verbose bool) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose
	err := http.ListenAndServe(addr, proxy)
	if err != nil {
		log.Fatal("Failed to start HTTP proxy: ", err)
	}
}

func makeListenAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
