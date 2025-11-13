package main

import (
	"fmt"
	"strconv"
)

// Parsing numbers from strings is a simple, but common task.
// The built-in package strconv provides the number parsing.

func main() {

	// With ParseFloat, the 64 tells the program how many bits of
	// precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// For ParseInt, the 0 means infer the base from the string,
	// with 64 enforcing that the result fit within 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt will recognise hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUInt is available for specifically unsigned numbers.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions will return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}