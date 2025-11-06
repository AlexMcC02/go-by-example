package main

import (
	"fmt"
	"time"
)

// A timeout is the amount of a time a program will wait for the completion of a certain event.

func main() {

	// Note that the channel here is buffered, meaning the send is non-blocking.
	// This is deliberate, and can prevent goroutine leaks if the channel is not read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// This select implements a timeout, with the first case awaiting the result
	// and the second implements the timeout, which is one second here.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Here we simulate c2 receiving its result in 2 seconds, with a timeout of 3 seconds.
	// Thus, it will succeed and print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}