package main

import (
	"fmt"
)

func main() { 
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // Control flow is simple in go, only notable aspect is that parentheses are not needed.
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Note: There is no ternary operator in Go.

	if num := 9; num < 0 { // The num variable is available in this and all subsequent branches.
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}