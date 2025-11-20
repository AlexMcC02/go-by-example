package main

import (
    "fmt"
    "os"
)

// Use os.Exit to immediately exit with a given status.

func main() {

    // defers will NOT run when using os.Exit, so this fmt.Println will never
    // be called.
    defer fmt.Println("!")

    // Unlike languages such as C, Go does not use an integer return value from
    // main to indicate exit status. To use a non-zero status, you should use 
    // os.Exit as demonstrated below.
    os.Exit(3)
}