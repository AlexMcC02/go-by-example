package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments, such as Println().

func sum(nums ...int) { // This function will take an arbitrary number of ints as arguments.
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums { // Within the function, the type of nums is equivalent to an []int.
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1 ,2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // You can pass a slice to a variadic function like this.
}