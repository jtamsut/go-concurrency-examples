package main

import (
	"fmt"
	"sync"
	"time"
)

// Container holds a map of counters that we will want to update
// concurrently from multiple goroutines

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {
	c := Container{counters: map[string]int{"slow": 0, "fast": 0}}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int, waitTime int) {
		time.Sleep(time.Duration(waitTime) * time.Second)
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("slow", 500, 2)
	go doIncrement("fast", 750, 1)
	go doIncrement("slow", 200, 2)

	wg.Wait()
	fmt.Println(c.counters)
}
