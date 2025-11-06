package main

import (
	"fmt"
	"time"
)

// Select allows you to combine goroutines and channels, it allows you to wait on multi-channel operations.

func main() {
	c1 := make(chan string) // The two channels we will be selecting across.
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time (simulating server-side operations).
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Select is used here to await for both of these values simultaneously, printing each one as it arrives.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}