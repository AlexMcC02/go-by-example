package main

import (
	"fmt"
	"slices"
)

func main() { 
	var s []string // Unlike arrays, slices are typed only by the elements they contain.
	fmt.Println("uninit", s, s == nil, len(s) == 0)
	// An unitialised slice equals to nil and has a length of 0.

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a" // Setting and getting is the same as arrays.
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // Returns the length of the slice as we'd expect.

	// Append is a unique feature of slices, which will return a slice containing one or more new values.
	s = append(s, "d") 
	s = append(s, "e", "f") // Note, we need to accept a return value as we may get a new slice value.
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) // Slices can be copied using the copy function.
	fmt.Println("cpy:", c)

	l := s[2:5] // Slice operator used to fetch elements between indices 2 and 5.
	fmt.Println("sl1:", l)

	l = s[:5] // Slices up until but excluding s[5].
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"} // One line declaration and initialisation.
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"} // Comparison of slices.
	if slices.Equal(t, t2) {
		fmt.Println("t == t2") // Slices are considered equal to their contents seemingly...?
	}

	// Slices can be composed into multi-dimensional data structures. The length o inner slices can vary
	// unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make ([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}