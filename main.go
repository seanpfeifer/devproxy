package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
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

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
