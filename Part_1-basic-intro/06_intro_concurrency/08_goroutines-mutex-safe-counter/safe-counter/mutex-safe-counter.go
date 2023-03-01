package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is a concurrent-safe map
type SafeCounter struct {
	mutex sync.Mutex
	m 		map[string]int
}

// Inc increments the value of the given key in the SafeCounter map
func (c *SafeCounter) Inc(key string) {
	c.mutex.Lock()
	c.m[key]++
	c.mutex.Unlock()
}

// Value returns the value associated to the given key in the SafeCounter map
func (c *SafeCounter) Value(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.m[key]
}

func main() {
	c := SafeCounter{m: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("someKey"))
}