package main

import "github.com/gronnbeck/latency/latency"

func main() {

	//latency.NewEtcdHTTPHandlerConfig("/test2", 0*time.Second, 2*time.Second)

	proxyURL := latency.ConfigProxyURL()

	config := latency.NewProbabilisticLatencyConfig(0, 3)
	proxy := latency.NewProxy(proxyURL, config)

	proxy.Listen(":8000")

}
