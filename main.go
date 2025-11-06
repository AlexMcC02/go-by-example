package main

import "fmt"

// Channels should be closed when there are no more values to be sent to it.
// This can communicate a completion state to the channel's receivers.

func main() {
	
	// A jobs channel will communicate the work to be completed by a worker goroutine.
	// When no more jobs need to be completed, the jobs channel will be closed.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// This is the worker goroutine, it repeatedly receives from jobs.
	// The more value indicates if there's more jobs to receive, if it
	// returns false, it means the jobs channel has been closed.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	// Reading from a closed channel succeeds immediately, returning the zero
	// value of the underlying type. The optional second return value is true if
	// the value received was delivered by a successful send operation to the
	// channel, or false if it was a zero value (generated because the channel is
	// both closed and empty).
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}