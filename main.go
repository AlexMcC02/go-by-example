package main

import "fmt"

// Standard send and receive operations on channels are blocking.
// Select, when used with a default clause, can implement non-blocking
// sends, receives, multi-way selects.

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// This is an example of a non-blocking receive.
	// If a value is available on the messages channel, select will take the appropriate case
	// otherwise, it will take the default case. This means something will always be returned.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send is similar. msg cannot be sent to the messages cahannel, as it
	// has no buffer and there's no receiver. The default case will be selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Multiple cases above the default case can be used to implement multi-way, 
	// non-blocking select.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}