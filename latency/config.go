package latency

import "os"

func init() {
	validateConfigProxyURL(os.Getenv("LATENCY_PROXY_URL"))

	if ConfigEtcdURL() == "" {
		panic("Missing envvar ETCD_URL")
	}
}

// ConfigProxyURL returns the URL latency is proxying to. Panics if
// envvar LATENCY_PROXY_URL is not set
func ConfigProxyURL() string {
	url := os.Getenv("LATENCY_PROXY_URL")
	return url
}

// ConfigEnvironment returns the current environment
func ConfigEnvironment() string {
	env := os.Getenv("ENVIRONMENT")

	return env
}

// ConfigEtcdURL returns the URL of the etcd to use
func ConfigEtcdURL() string {
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
