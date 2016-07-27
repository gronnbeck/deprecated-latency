package config

import "os"

func init() {
	validateConfigProxyURL(os.Getenv("LATENCY_PROXY_URL"))

	if EtcdURL() == "" {
		panic("Missing envvar ETCD_URL")
	}
}

// ProxyURL returns the URL latency is proxying to. Panics if
// envvar LATENCY_PROXY_URL is not set
func ProxyURL() string {
	url := os.Getenv("LATENCY_PROXY_URL")
	return url
}

// Environment returns the current environment
func Environment() string {
	env := os.Getenv("ENVIRONMENT")

	return env
}

// EtcdURL returns the URL of the etcd to use
func EtcdURL() string {
	return os.Getenv("ETCD_URL")
}

func validateConfigProxyURL(url string) {
	if url == "" {
		panic("Missing envvar LATENCY_PROXY_URL")
	}

	if url[0:7] != "http://" && url[0:8] != "https://" {
		panic("Unsupported protocol. Only support http:// or https://")
	}
}
