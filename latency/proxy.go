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

// NewHTTPHandler creates and returns a new HTTPHandler
func NewHTTPHandler(url string, config HTTPHandlerConfig) HTTPHandler {
	return HTTPHandler{proxyURL: url, config: config}
}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	config := h.config

	req2 := h.copyRequest(req)

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

		if getHeaderInfoEnabled() {
			w.Header().Add("X-LATENCY-ADDED-LATENCY", latency.String())
		}

	}

	w.WriteHeader(res.StatusCode)
	w.Write(byt)
}

func (h HTTPHandler) copyRequest(req *http.Request) *http.Request {
	copy, err := http.NewRequest(req.Method, h.proxyURL, req.Body)
	copy.Header = req.Header

	if err != nil {
		panic(err)
	}

	return copy
}

func getHeaderInfoEnabled() bool {
	return ConfigEnvironment() == "development"
}
