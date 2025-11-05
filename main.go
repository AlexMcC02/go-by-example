package main

// A goroutine is a lightweight thread of execution.

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct") // This is the usual way we'd call a function -- it is run synchronously.
	go f("goroutine") // This is how we invoke a goroutine, it will execute concurrently with the calling one.
	go func(msg string) { // You can also start a goroutine for an anonymous function call.
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) // The two function calls are running asynchronously in separate goroutines. Here we wait for them to finish.
	fmt.Println("done") // The synchronous function will always execute first, as it synchronous execution is 'blocking'.
}