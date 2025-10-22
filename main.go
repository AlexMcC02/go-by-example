package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	// Range iterates over elements in a variety of built-in data structures.

	for _, num := range nums { // Range is used here to summate the elements in a slice.
		sum += num
	}
	fmt.Println("sum:", sum)

	// Note, using range will provide BOTH the index and value.
	// You will often see one of these discarded, as is exemplified below.
	for i, num := range nums { 
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on a map will iterate over the key-value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Though it can also just iterate over the keys.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Using range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the ruie and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}