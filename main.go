package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 { // Simple for loop with a single condition (this would be a while loop in other languages).
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // The classic initial;condition;after for loop.
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i) // Accomplishes the same as the above, as i's default value will be 0.
	}

	for { // Will loop endlessly until broken (again, much like a While true statement).
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue // You can use the continue keyboard to move to the next iteration of the loop.
		}
		fmt.Println(n)
	}
}