package main

import (
    "fmt"
    "os"
)

// Defer is used to ensure that a function call is preformed later in a program's execution.
// This is usually done for the purposes of cleanup. Defer is usually done in place of a 
// 'finally' block in other languages.

func main() {

	// Below is an example of creating, writing, and closing a file whilst taking advantage of defer.
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

// It is still important to check for errors when closing a file, even if it is a deferred function.
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        panic(err)
    }
}