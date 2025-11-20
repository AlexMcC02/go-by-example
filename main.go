package main

import (
    "bufio"
    "fmt"
    "net/http"
)

// The Go standard library comes with excellent support for HTTP clients and 
// servers in the net/http package.

func main() {

    // Here we issue an HTTP GET request to a server. http.Get is a convenient
    // shortcut around creating an http.Client object and calling its Get method,
    // with it invoking a http.DefaultClient object that contains sensible defaults.
    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    // Here we output the first five lines of the response body.
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}