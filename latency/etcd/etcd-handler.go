package etcd

import (
	"sync"
	"time"
)

// Handler listens to changes in etcd and updates it's values accordingly.
// It is safe for concurrent reads and updates.
type Handler struct {
	mu  *sync.RWMutex
	min time.Duration
	max time.Duration
}

// NewHandler creates a Handler with the specified min and max values.
func NewHandler(min, max time.Duration) *Handler {
	return &Handler{
		mu:  new(sync.RWMutex),
		min: min,
		max: max,
	}
}

// GetMin returns min.
func (h Handler) GetMin() time.Duration {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.min
}

// SetMin sets min to specified value
func (h *Handler) SetMin(min time.Duration) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.min = min
}

// GetMax returns max
func (h Handler) GetMax() time.Duration {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.max
}

// SetMax sets max to specified value
func (h *Handler) SetMax(max time.Duration) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.max = max
}
