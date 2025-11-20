package main

import (
    "fmt"
    "net/http"
    "time"
)

// In the prior example we looked at setting up a simple HTTP server. HTTP
// servers are useful for demonstrating the usage of context.Context for 
// controlling cancellation. A context carries deadlines, cancellation
// signals, and other request-scoped values across API boundaries and
// goroutines.

func hello(w http.ResponseWriter, req *http.Request) {

    // A context.Context is created for each requests by the net/http machinery,
    // and is available with the Context() method.
    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    // Here we simulate work with a delay. We also keep an eye on the context's
    // Done() channel for a singla that we should cancel the work and return immediately.
    select {
    case <-time.After(10 * time.Second):
        fmt.Fprint(w, "hello\n")
    case <-ctx.Done():

        // The context's Err() method returns an error that explains why the Done()
        // channel was closed.
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}