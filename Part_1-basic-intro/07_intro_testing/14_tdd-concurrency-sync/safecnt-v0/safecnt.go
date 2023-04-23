package safecnt

// Counter is a concurrent-safe counter
type Counter struct {
	value int
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.value++
}

// Value lets you interrogate the current value of the counter
func (c *Counter) Value() int {
	return c.value
}