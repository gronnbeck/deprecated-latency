package latency

import (
	"sync"
	"time"
)

// Store listens to changes in etcd and updates it's values accordingly.
// It is safe for concurrent reads and updates.
type Store struct {
	mu  *sync.RWMutex
	min time.Duration
	max time.Duration
}

// NewStore creates a Store with the specified min and max values.
func NewStore(min, max time.Duration) *Store {
	return &Store{
		mu:  new(sync.RWMutex),
		min: min,
		max: max,
	}
}

// GetMin returns min.
func (h Store) GetMin() time.Duration {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.min
}

// SetMin sets min to specified value
func (h *Store) SetMin(min time.Duration) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.min = min
}

// GetMax returns max
func (h Store) GetMax() time.Duration {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.max
}

// SetMax sets max to specified value
func (h *Store) SetMax(max time.Duration) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.max = max
}
