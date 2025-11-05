package main

// The idomatic way to communicate errors in Go is with an explicit, separate return type.
// This may be quite strange if you're used to Python, Java, and Ruby, which make use of exceptions.
// Go's approach is intended to make it easy to see which functions return errors and to handle them
// using the same language constructs employed for other, non-error tasks.

import (
	"fmt"
	"errors"
)

// By convention, errors are the last erturn value and have a type error, a built-in form of type inference.
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42") // errors.New constructs a basic error value with the supplied message.
	}
	return arg + 3, nil // In this case, a return value of nil signals that there was no error.
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

// We can wrap errors with higher-level errors to add context. The %w verb in fmt.Errorf can be used to achieve this.
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("make tea: %w", ErrPower)
	}
	return nil
}

// It's idiomatic to use an inline error check in the if line.
func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	// errors.Is checks that a given error matches a specific error value.
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy never tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		
		fmt.Println("Tea is ready!")
	}
}