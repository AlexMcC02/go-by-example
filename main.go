package main

import "fmt"

// When providing channels as function parameters, you can specify if it should only send or receive values.
// This specificity is intended to increase type safety.

// This function only accepts a channel for sending values. 
// Attempting to receive on this channel will result in a compile time error.
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// This function accepts one channel for receivers (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}