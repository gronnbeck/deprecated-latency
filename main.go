package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gronnbeck/latency/latency"
)

func main() {

	proxyURL := latency.ConfigProxyURL()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		req2, err := http.NewRequest(req.Method, proxyURL, req.Body)
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
	})

	http.ListenAndServe(":8000", nil)

}
