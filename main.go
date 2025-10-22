package main

import "fmt"

func fact(n int) int { // This function will call itself until it reaches the base case of fact(0).
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(5))

	// Anonymous functions can also be recursive, but you must explicitly declare a variable with var
	// to store the function before it's defined.
	var fib func(n int) int 

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}