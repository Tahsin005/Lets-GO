package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v map[string]int
	mu sync.Mutex
}

func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	// lock, so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increment("a")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("a"))
}
