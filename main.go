package main

import (
	"net/http"

	"github.com/gronnbeck/latency/latency"
)

func main() {

	proxyURL := latency.ConfigProxyURL()

	config := latency.NewProbabilisticLatencyConfig(0, 3)
	latencyHandler := latency.NewHTTPHandler(proxyURL, config)

	http.Handle("/", latencyHandler)

	http.ListenAndServe(":8000", nil)

}
