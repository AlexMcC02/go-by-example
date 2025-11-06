package main

import (
    "fmt"
    "sync"
    "time"
)

// When waiting for multiple goroutines to finish, we can use a wait group.

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This wait group is used to wait for all the goroutines launched here to finish.
	// Note: if a wait group is explicitly passed into a function, a pointer should be used.
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Go(func() {
            worker(i)
        })
    }

	// Here we use the wait group to block all goroutines until they're done.
	// A goroutine is considered 'done' when the function it invokes returns.
    wg.Wait()

}