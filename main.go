package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// A Go string is read-only sequence of bytes.
	// The language and standard library treat strings as containers of text encoded in UTF-8.
	// In other languages, strings are made of "characters".
	// In Go, the concept of a character is called a Rune, an integer that represents a Unicode code point.

	// S is a string assigned a literal value representing the word "hello" in the Thai language.
	// Go string literals are UTF-8 encoded text.
	const s = "สวัสดี"

	// Because strings are equivalent to []byte, this statement will produce the length of the 
	// raw bytes stored within.
	fmt.Println("Len:", len(s))

	// Indexing a string produces the raw byte values at each index. This loop will output the
	// hexidecimal values of all the bytes that constitute the code points in s.
	for i := range len(s) {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()

	// To count the number of runes in as tring, we use the utf8 package. Because some Thai characters
	// are represented by UTF-8 code points that can span multiple bytes, the result of this count may
	// be larger than you expected.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings in a special way, decoding each rune with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// The same results can be achieved using the utf8.DecodeRuneInString function.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Passing a rune to a function.
		examineRune(runeValue)
	}
}

// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}