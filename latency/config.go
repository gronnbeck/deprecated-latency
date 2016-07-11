package latency

import "os"

// ConfigProxyURL returns the URL latency is proxying to. Panics if
// envvar LATENCY_PROXY_URL is not set
func ConfigProxyURL() string {
	url := os.Getenv("LATENCY_PROXY_URL")

	validateConfigProxyURL(url)

	return url
}

// ConfigEnvironment returns the current environment
func ConfigEnvironment() string {
	env := os.Getenv("ENVIRONMENT")

	return env
}

func validateConfigProxyURL(url string) {
	if url == "" {
		panic("Missing envvar LATENCY_PROXY_URL")
	}

	if url[0:7] != "http://" && url[0:8] != "https://" {
		panic("Unsupported protocol. Only support http:// or https://")
	}
}
