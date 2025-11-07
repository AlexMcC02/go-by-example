package main

import (
	"cmp"
	"fmt"
	"slices"
)

// There are times you'll want to sort a collection by something other than its natural order.
// E.g. -- we want to sort strings by their length instead of alphabetically.

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string lengths (cmp.Compare is helpful for this).
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Now we can call slices.SortFunc with this custom comparison to sort fruits by name length.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// The same technique can be employed to sort a slice of custom values.
	type Person struct {
		name string
		age int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// This implementation sorts the custom type of Person by their age.
	// For performance reasons, it may be desirable to pass my reference rather than value with a pointer.
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}