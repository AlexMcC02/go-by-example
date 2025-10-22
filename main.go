package main

import (
	"fmt"
)

func plus(a int, b int) int { // Return type after the parameters.
	return a + b // Explicit turns, hallejulah.
}

func plusPlus(a, b, c int) int { // Declare the type of multiple parameters like this. Very cool!
	return a + b + c
}

func main() { 
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}