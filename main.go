package main

import (
	"fmt"
	"maps"
)

func main() { 
	m := make(map[string]int) // String key and an int value.

	m["k1"] = 7 // Setting key value pairs with a standard [key] = value syntax.
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1) // Get a specific value based on a key.

	v3 := m["k3"]
	fmt.Println("v3:", v3) // A zero value is returned if the key does not exist.

	fmt.Println("len:", len(m)) // Returns the number of key value pairs.

	delete(m, "k2") // Delete keyword does, you guessed it, delete things! O.o
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m) // Removes all key-value pairs from a map.

	_, prs := m["k2"] // Optional return value indicates if a key is present in the map. _ used to discard the value itself.
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // Compact init and declare syntax for maps.
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) { // Like with slices, we can compare maps.
		fmt.Println("n == n2")
	}
}