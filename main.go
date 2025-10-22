package main

import "fmt"

// Go supports anonymous functions, which can form closures. These are useful when you want to define an
// inline function without having to name it.

// A closure retains the state of the variables accessed from its surrounding scope (i.e. the function the closure)
// is nested inside.

// This function returns another function, which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form a closure.
func intSeq() func() int { 
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	// Each time intSeq is called, we assign it to nextInt. Because the function value captures 
	// its own i value, it wil be updated each time we call nextInt.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// If we create a new variable and assign intSeq to it, we get a unique, unrelated value.
	newInts := intSeq()
	fmt.Println(newInts())
}