package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

// There are times when you want a Go program to intelligently handle Unix signals.

func main() {

    // Go signal notification works by sending os.Signal values on a chnnel. Note that
    // the channel we've created here is buffered.
    sigs := make(chan os.Signal, 1)

    // signal.Notify registers the given channel to receive notifications of the
    // specified signals.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Below demonstrates how receives can be received from a separate goroutine.
    done := make(chan bool, 1)

    // This goroutine executes a blocking receive for signals. When it gets one
    // it'll print it out then notify the program it can finish.
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    // The program will wait until it gets the expected signal before exiting.
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}