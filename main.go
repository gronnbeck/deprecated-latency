package main

import (
	"net/http"

	"github.com/gronnbeck/latency/latency"
)

func main() {

	proxyURL := latency.ConfigProxyURL()

	latencyHandler := latency.NewHTTPHandler(proxyURL)

	http.Handle("/", latencyHandler)

	http.ListenAndServe(":8000", nil)

}
