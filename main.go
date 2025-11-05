package main

import (
    "fmt"
    "time"
)

// Channels can be used to synchronise execution across goroutines.

// In this function, the done channel will be used to notify another goroutine that this function's work is done.
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true // Sending a value to notify that the function is done.
}

func main() {

    done := make(chan bool, 1) // Starting a worker goroutine, giving it a channel to notify on.
    go worker(done) 

    <-done // Block until we receive a notification from the worker on the channel.
}