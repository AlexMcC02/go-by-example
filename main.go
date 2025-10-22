package main

import (
	"fmt"
)

func main() { 
	var a [5]int
	fmt.Println("emp:", a) // Zero valued by default [0 0 0 0 0].

	a[4] = 100 // Set the fifth value to be 100.
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) // Retrieve a specific element by its index.

	fmt.Println("len:", len(a)) // Return the length of the array.

	b := [5]int{1, 2, 3, 4, 5} // One line array initialisation and declaration.
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} // Using [...] will have the compiler apply an appropriate length for you.
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} // You can use the : symbol to specify a starting index.
	fmt.Println("idx:", b)

	var twoD [2][3]int // You can compose types to build multi-dimensional data structures.
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int { // Initialising a multi-dimensional array.
		{1, 2, 3},
		{4, 5 ,6},
	}
	fmt.Println("2d:", twoD)
}