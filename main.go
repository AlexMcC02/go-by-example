package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// Communication over channels is the primary means of state management. This was demonstrated with
// worker pools. Another means to achieve this is through atomic counters.

func main() {

	// Declaring an atomic integer type.
    var ops atomic.Uint64

	// This wait group will help us wait for all goroutines to finish their work.
    var wg sync.WaitGroup

	// Here we start 50 goroutines that each increment the counter 1000 times.
    for range 50 {
        wg.Go(func() {
            for range 1000 {

                ops.Add(1)
            }
        })
    }

	// Call the wait group to block until all goroutines are done.
    wg.Wait()

	// Here no goroutines are writing to ops, but using Load it's safe to atomically read
	// a value even while other goroutines are atomically updating it.
    fmt.Println("ops:", ops.Load())
}