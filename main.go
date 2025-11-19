package main

import (
    "fmt"
    "flag"
)

// Command-line flags are a common means to specify options for CLI programs.
// The flag package provides support for command line parsing.

func main() {

    // Basic flag decorations are available for string, integer, and boolean
    // options. Note that the flag.String function will return a pointer,
    // not a value.
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    // It's also possible to declare an option that uses an existing var 
    // declared elsewhere program. Note that we need to pass a pointer to
    // the flag declaration function.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Flag parse is then called to execute the command-line parsing.
    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}