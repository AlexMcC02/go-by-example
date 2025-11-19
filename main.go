package main

import (
    "fmt"
    "os"
)

// Command-line arguments are a common way to parameterize execution of programs.

func main() {

    // os.Args provide access to raw command-line arguments. Note that the first
    // value in this slice is the path to the program, with os.Args[1:] holds the
    // arguments to the program.
    argswithProg := os.Args
    argsWithoutProg := os.Args[1:]

    // Arguments can be accessed using an index.
    arg := os.Args[3]

    fmt.Println(argswithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}