package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

// A line filter is a common type of program that reads input on stdin, process it,
// and then prints some derived result to stdout. grep and sed are common line filters.

// This example has a line filter that writes a capitalised version of provided input text.
func main() {

	// Wrapping the unbuffered os.Stdin with a buffered scanner provides a convenient 
	// Scan method that advances the scanner to the next token (the nextg line in the
	// default scanner).
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Error check that gracefully exits the application.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}