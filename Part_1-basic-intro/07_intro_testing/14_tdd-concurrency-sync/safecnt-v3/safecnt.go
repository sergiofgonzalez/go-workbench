package safecnt

import "sync"

// Counter is a concurrent-safe counter
type Counter struct {
	mx    sync.Mutex
	value int
}

// NewCounter returns a pointer to a properly initialized counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.value++
}

// Value lets you interrogate the current value of the counter
func (c *Counter) Value() int {
	return c.value
}
