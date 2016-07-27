package latency

import "time"

// HTTPHandlerConfig describes how a HTTPHandler should behave
type HTTPHandlerConfig interface {
	GetLatency() *time.Duration
}
