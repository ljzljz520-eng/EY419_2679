package audit

import (
	"sync"
	"time"
)

type Counter struct {
	mu      sync.Mutex
	values  map[string]int64
	updated time.Time
}

func NewCounter() *Counter { return &Counter{values: map[string]int64{}, updated: time.Now().UTC()} }
func (c *Counter) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[name]++
	c.updated = time.Now().UTC()
}
func (c *Counter) Add(name string, n int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[name] += n
	c.updated = time.Now().UTC()
}
func (c *Counter) Get(name string) int64 { c.mu.Lock(); defer c.mu.Unlock(); return c.values[name] }
func (c *Counter) Snapshot() map[string]int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := map[string]int64{}
	for k, v := range c.values {
		out[k] = v
	}
	return out
}
func (c *Counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values = map[string]int64{}
	c.updated = time.Now().UTC()
}
