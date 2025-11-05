package main

// In Go, it is possible to define custom error types by implementing the Error() method on them.

import (
	"errors"
	"fmt"
)

type argError struct { // A custom error type will usually have the suffix "Error".
	arg int
	message string
}

func (e *argError) Error() string { // Adding this Error method makes argError implement the error interface.
	return fmt.Sprint("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} // Returning the custom error.
	}
	return arg + 3, nil
}

// errors.As is a more advanced version of errors.Is. It checks that a given error (or any error in its chain) matches a specific error
// type and converts to a value of that type, returning true. If there's no match, it will return false.
func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}