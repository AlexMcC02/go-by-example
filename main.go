package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go offers built-in support for regular expressions.

func main() {
	
	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Above a string pattern was used directly, but for other regexp tasks you'll
	// need to compile an optimised Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Matching for the presence of a string.
	fmt.Println(r.MatchString("peach"))
	
	// Finds the match for the regexp.
	fmt.Println(r.FindString("peach punch"))

	// Returns the start and indexes for the match instead of the matching text.
	fmt.Println("idx:", r.FindStringIndex("peach peach"))

	// This function will include information about both the whole-pattern matches
	// and the submatches within those matches.
	fmt.Println(r.FindStringSubmatch("peach peach"))

	// This will return the indexes of the matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The all variants of therse functions to apply to all matches
	// in the input, not just the first.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Similar case as above.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

	// Supplying a non-negative second argument will limit the 
	// number of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// In addition to strings, []byte arguments can be supplied instead.
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile makes it safer to make use of global variables for regular
	// expressions, as it will panic instead of returning an error.
	r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)

	// The regexp package also enables the replacing of string subsets.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Finally, the Func variant allows you to transform the matched
	// text within a given fujnction.
	in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}