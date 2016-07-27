package latency

import (
	"math/rand"
	"time"
)

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

// ProbabilisticLatencyConfig controls latency using a probabilistic model
type ProbabilisticLatencyConfig struct {
	minLatency int
	maxLatency int
}

// NewProbabilisticLatencyConfig returns a freshly created ProbabilisticLatencyConfig
func NewProbabilisticLatencyConfig(minLatency, maxLatency int) ProbabilisticLatencyConfig {
	return ProbabilisticLatencyConfig{minLatency, maxLatency}
}

// GetLatency returns the latency based on configuration of the
// ProbabilisticLatencyConfig
func (c ProbabilisticLatencyConfig) GetLatency() *time.Duration {
	value := c.minLatency + rand.Intn(c.maxLatency-c.minLatency)
	latency := time.Duration(value) * time.Second
	return &latency
}
