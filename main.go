package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

// Go offers a variety of string formatting options with the printf function.

func main() {

	// Go offers several printing 'verbs' designed to format general Go values.
	// E.g. the %v verb prints an instance of the point struct.
    p := point{1, 2}
    fmt.Printf("struct1: %v\n", p)

	// The %v+ variant includes the field names of the struct.
    fmt.Printf("struct2: %+v\n", p)

	// The %#v variant includes a Go syntax representation of the value or
	// the source code snippet that would produce the value.
    fmt.Printf("struct3: %#v\n", p)

	// %T prints the type of the value.
    fmt.Printf("type: %T\n", p)

	// Example of formatting booleans.
    fmt.Printf("bool: %t\n", true)

	// %d is used for standard base 10 formatting.
    fmt.Printf("int: %d\n", 123)

	// %b prints a binary representation.
    fmt.Printf("bin: %b\n", 14)

	// %c prints the character corresponding to the given integer.
    fmt.Printf("char: %c\n", 33)

	// %x provides hex encoding.
    fmt.Printf("hex: %x\n", 456)

	// %f is used for standard floating point values.
    fmt.Printf("float1: %f\n", 78.9)

	// %e and %E format the float using scientific notation.
    fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)

	// %s is used for basic string formatting.
    fmt.Printf("str1: %s\n", "\"string\"")

	// %q can be used to allow the use of double quotes in strings.
    fmt.Printf("str2: %q\n", "\"string\"")

	// %x can also be used to render strings in base 16.
    fmt.Printf("str3: %x\n", "hex this")

	// %p can be used to print a representation of a pointer.
    fmt.Printf("pointer: %p\n", &p)

	// You can specify the width of a printed figure by placing a number before %d.
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// You can apply the width.precision syntax as shown below for floats.
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// By default numbers will be right justified, you can apply - flag to left justify instead.
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting strings.
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// Again the - flag is used to left justify.
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without printing it.
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

	// You can format and print to io.Writers other than os.Stdout using Fprintf.
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}