package main

import "fmt"

func vals() (int, int) { // This function returns two integer values.
	return 3, 7
}

// Handling multiple returns is common in Go, with many functions returning a result and error value.

func main() {
	a, b := vals() // Multiple assignments work well with functions returning multiple values.
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // If we only want one of the values, we can opt to discard a value with _.
	fmt.Println(c)
}