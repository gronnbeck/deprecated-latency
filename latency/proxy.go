package latency

import (
	"io/ioutil"
	"net/http"
)

// HTTPHandler introduces latency to a http handler
type HTTPHandler struct {
	proxyURL string
}

// NewHTTPHandler creates and returns a new HTTPHandler
func NewHTTPHandler(url string) HTTPHandler {
	return HTTPHandler{proxyURL: url}
}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

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

	w.Write(byt)
}
