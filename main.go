package main

import "fmt"

// By default, channels are unbuffered.
// This means that channels will only accept sends (chan <-) if there is a corresponding receive (<- chan).
// Buffered channels on the other hand accept a limited number of values without a corresponding receiver for said values.

func main() {

	// A channel of strings buffering up to two values.
    messages := make(chan string, 2)

	// As the channel is buffered, we don't need to a corresponding concurrent receive.
    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}