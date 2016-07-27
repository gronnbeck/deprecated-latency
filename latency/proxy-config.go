package latency

import "time"

// ProxyConfig describes how a HTTPHandler should behave
type ProxyConfig interface {
	GetLatency() *time.Duration
}
