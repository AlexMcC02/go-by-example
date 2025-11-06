package main

import (
	"fmt"
	"time"
)

// Tickers are useful for when you want to do something repeatedly at regular intervals.

func main() {

	// Tickers use a similar mechanism to timers: a channel that is sent values.
	// Here we'll use the select builtin on the channel to await values as they
	// arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Like timers, tickers can be stopped. When a ticker is stopped it won't receive
	// any more values on its associated channel.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}