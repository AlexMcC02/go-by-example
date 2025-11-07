package main

import (
	"fmt"
	"sync"
)

// To manage complex state changes, we can use a mutex to safely access data across multiple goroutines.

// Container holds a map of counters. Because we want to update it concurrently from multiple goroutines
// a mutex is added to synchronise access. 
// Note: mutexes must not be copied, thus if a struct containing one is passed around it, it must be done
// with a pointer.
type Container struct {
	mu sync.Mutex
	counters map [string]int
}

// Here we lock a mutex before accessing counters. The defer statement will then unlock it.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock() // Defer ensures a function call is performed later in a program's execution.
	c.counters[name]++
}

// Note: a mutex's zero value can be used as-is, and doesn't require explicit initialisation.
func main() {
	c := Container {
		counters: map[string]int{"a":0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Several goroutines are ran concurrently below, with them all accessing the same container,
	// with two accessing the same counter.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	wg.Wait()
	fmt.Println(c.counters)
}