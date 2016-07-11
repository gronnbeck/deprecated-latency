package latency

import (
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPHandler introduces latency to a http handler
type HTTPHandler struct {
	proxyURL string
	config   HTTPHandlerConfig
}

// HTTPHandlerConfig describes how a HTTPHandler should behave
type HTTPHandlerConfig interface {
	GetLatency() *time.Duration
}

// NewHTTPHandler creates and returns a new HTTPHandler
func NewHTTPHandler(url string, config HTTPHandlerConfig) HTTPHandler {
	return HTTPHandler{proxyURL: url, config: config}
}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	config := h.config

	req2, err := http.NewRequest(req.Method, h.proxyURL, req.Body)
	req2.Header = req.Header

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req2)

	if err != nil {
		panic(err)
	}

	byt, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	if config.GetLatency() != nil {
		latency := *config.GetLatency()
		time.Sleep(latency)

		w.Header().Add("X-LATENCY-ADDED-LATENCY", latency.String())
	}

	w.WriteHeader(res.StatusCode)
	w.Write(byt)
}

// FixedLatencyConfig introduces the same latency for each request
type FixedLatencyConfig struct {
	latency time.Duration
}

// NewFixedLatencyConfig returns a fresh FixedLatencyConfig
func NewFixedLatencyConfig(latency time.Duration) FixedLatencyConfig {
	return FixedLatencyConfig{latency}
}

// GetLatency returns the same latency defined in its construction
func (c FixedLatencyConfig) GetLatency() *time.Duration {
	return &c.latency
}
