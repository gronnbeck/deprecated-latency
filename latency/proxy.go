package latency

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gronnbeck/latency/config"
)

// Proxy introduces latency to a http handler
type Proxy struct {
	proxyURL string
	config   ProxyConfig
}

// NewProxy creates and returns a new Proxy
func NewProxy(url string, config ProxyConfig) Proxy {
	return Proxy{proxyURL: url, config: config}
}

// Listen is a shorthand method for setting up a Proxy.
// Alternatively you can use Proxy and set it up manually using http.Handle
func (h Proxy) Listen(uri string) {
	http.Handle("/", h)
	http.ListenAndServe(uri, nil)
}

func (h Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

func (h Proxy) copyRequest(req *http.Request) *http.Request {
	copy, err := http.NewRequest(req.Method, h.proxyURL, req.Body)
	copy.Header = req.Header

	if err != nil {
		panic(err)
	}

	return copy
}

func getHeaderInfoEnabled() bool {
	return config.Environment() == "development"
}
