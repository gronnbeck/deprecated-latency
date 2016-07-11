package main

import (
	"net/http"
	"time"

	"github.com/gronnbeck/latency/latency"
)

func main() {

	proxyURL := latency.ConfigProxyURL()

	config := latency.NewFixedLatencyConfig(3 * time.Second)
	latencyHandler := latency.NewHTTPHandler(proxyURL, config)

	http.Handle("/", latencyHandler)

	http.ListenAndServe(":8000", nil)

}
