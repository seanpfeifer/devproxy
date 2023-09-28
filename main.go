package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	defaultPort = 8080
)

func main() {
	port := flag.Int("port", defaultPort, "the port the reverse proxy will listen on")
	allTargets := proxyTargets{}
	flag.Var(&allTargets, "proxy", `a string in the format of "/url/on/local/->http://remote:port/url/on/remote/"`)
	flag.Parse()

	if len(allTargets) == 0 {
		log.Fatal("No proxy targets specified.")
	}

	mux := http.NewServeMux()

	for _, target := range allTargets {
		local, remote, err := target.Parse()
		if err != nil {
			log.Fatalf(`Failed to parse proxy target "%s": %v`, target, err)
		}

		log.Printf(`Proxying "%s" -> "%s"`, local, remote)
		mux.Handle(local, httputil.NewSingleHostReverseProxy(remote))
	}

	hostAddr := fmt.Sprintf(":%d", *port)
	log.Printf(`Listening on "%s"`, hostAddr)

	err := http.ListenAndServe(hostAddr, mux)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
