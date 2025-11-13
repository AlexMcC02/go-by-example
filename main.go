package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
)

// Writing files in go follows similar patterns to the one we saw
// earlier for reading.

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	// To start, this is how to dump a string into a file.
    d1 := []byte("hello\ngo\n")
    path1 := filepath.Join(os.TempDir(), "dat1")
    err := os.WriteFile(path1, d1, 0644)
    check(err)

	// For more granular writes, open a file for writing.
    path2 := filepath.Join(os.TempDir(), "dat2")
    f, err := os.Create(path2)
    check(err)

	// It is idiomatic to close immediately after opening, using defer.
    defer f.Close()

	// Writing byte slices is quite straightforward.
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

	// A WriteString is also available.
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
    f.Sync()

	// Bufio provides buffered writers in addition to the buffered readers we
	// saw earlier.
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

	// Flush ensures all buffered operations have been applied to the underlying
	// writer.
    w.Flush()

}