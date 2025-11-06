package main

import (
    "fmt"
    "time"
)

// In this example, we implement a worker pool using goroutines and channels.

// This is the worker, of which there'll be several instances. Workers receive
// work on the jobs channel and send the corresponding results on results. We'll
// sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

	// We'll make use of a jobs channel and a results channel to send work to the
	// workers and collect their results.
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

	// After all five jobs are sent, the channel is closed.
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

	// Finally, the results of the work are collected. This will also ensure that worker
	// goroutines have finished. In the next example, we'll demonstrate a similar
	// implementation using work groups.
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}