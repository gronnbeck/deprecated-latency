package latency

import "testing"

func Test_validateConfigProxyURL(t *testing.T) {
	validateConfigProxyURL("http://thisurl")
	validateConfigProxyURL("https://thisurl")
}
