package main

import "os"

// A panic typically means something went unexpectedly wrong. It is most useful to 'fail fast'
// on errors that should NOT occur during normal operation, or that we haven't prepared to handle
// in a more graceful way (by throwing and catching).

func main() {
	panic("a problem") // When this panic fires, the other code will not be reached.

	_, err := os.Create("tmp/file")
	if err != nil {
		panic(err)
	}
}