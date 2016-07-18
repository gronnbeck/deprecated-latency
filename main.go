package main

import (
	"net/http"
	"time"

	"github.com/gronnbeck/latency/latency"
)

func main() {

	latency.NewEtcdHTTPHandlerConfig("/test2", 0*time.Second, 2*time.Second)

	proxyURL := latency.ConfigProxyURL()

	config := latency.NewProbabilisticLatencyConfig(0, 3)
	latencyHandler := latency.NewHTTPHandler(proxyURL, config)

	http.Handle("/", latencyHandler)

	http.ListenAndServe(":8000", nil)

}
