package main

import (
    "fmt"
    "os"
    "strings"
)

// Environment variables are a universal mechanism for conveying configuration
// information to Unix programs.

func main() {

    // os.Setenv is used to set key/value pairs.
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR")) // This will return an empty string.

    fmt.Println()

    // os.Environ is used to obtain a list of all key/value pairs in the environment.
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}