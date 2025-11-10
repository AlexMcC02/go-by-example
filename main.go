package main

import "fmt"

// Recover is a built-in function that allows you to recover from a panic.
// A recover can stop a panic from aborting the program and let it continue with execution instead.

// This function will panic.
func mayPanic() {
	panic("a problem")
}

func main() {
	// Recover must be called within a deferred function.
	// When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// This code won't run as mayPanic() will panic.
	// The execution of main will cease at this point and resume in the deferred closure.
	fmt.Println("After myPanic()")
}