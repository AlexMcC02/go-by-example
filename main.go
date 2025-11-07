package main

import (
	"fmt"
	"time"
)

// Rate limiting is a key mechanism for controlling resource utilisation and maintaining quality of service.
// Go elegantly supports rate limiting with goroutines, channels, and tickers.

func main() {
	// Below is an example of basic rate limiting. For example, if we want to limit how incoming requests are handled.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// The limited channel receives a value every 200 milliseconds, this will serve to regulate the rate limit.
	limiter := time.Tick(200 * time.Millisecond)

	// The limiter is implemented by blocking on a receive from it, this will limit the application to processing
	// 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// It may be desirable to have short bursts of requests withion the rate limiting scheme, whilst preserving
	// the overall limit. This can be accomplished by buffering the limiter channel. This channel will allow bursts
	// of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Simulating filling the buffer.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// In this example, the program will attempt to add a new value to the burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// In this example, 5 more incoming requests are simulated. The first 3 will benefit from the capacity of
	// the burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}