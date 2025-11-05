package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines. You can send values 
// into channels from one goroutine and receive those values into another goroutine.

func main() {

	// New channels are created with make, with them being typed by the values they convey.
    messages := make(chan string)

	// Here a value is sent into a channel using the <- syntax.
    go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel.
    msg := <-messages

	// Upon running the program, the 'ping' message will be passed from one goroutine to another via the messages channel.
	// By default send and receives block until both sender and receiver a ready.
    fmt.Println(msg)
}