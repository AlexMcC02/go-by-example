package main

import "fmt"

// Prior examples have made use of for and range to provide iteration over basic data structures.
// The same syntax can be used to iterate over values received from a channel.

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

	// This range iterates over each element as it's received from the queue channel.
	// Because the above channel has been closed, the iteration will terminate after
	// receiving the 2 elements.
    for elem := range queue {
        fmt.Println(elem)
    }

	// As you can see, it's possible to a close a non-empty channel but still have the
	// remaining values be received.
}