package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// Another option to synchronise shared state across goroutines, is to make use of the built-in
// features of goroutines and channels to achive a similar result. This approach aligns with Go's
// ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.

// In this example, our state will only ever be owned by a single goroutine. This ensures data is never
// corrupted with concurrent access. Instead, if other goroutines seek access, they achieve this by
// sending and receiving messages from the current, owning goroutine.

// The readOp and writeOp structs below encapsulate these requests and provide a way for the owning
// goroutine to respond.
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

	// Variables to keep track of the number of operations performed.
    var readOps uint64
    var writeOps uint64

	// The reads and writes channels will be used by other goroutines
	// to issue read and write requests.
    reads := make(chan readOp)
    writes := make(chan writeOp)

	// This will be the goroutine that owns the state.

	// This goroutine will repeatedly select on the reads and writes
	// channels, responding to requests as they arrive.

	// A response is executed by performing the request operation and then
	// sending a value on the response channel to indicate success as well as the
	// desired value in the case of reads.
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

	// This loop starts up 100 goroutines to issue reads to the state-owning goroutine via
	// the reads channel. Each read requires constructing a readOp, sending it over the reads
	// channel, and then receiving the result over the provided response channel.
    for range 100 {
        go func() {
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for range 10 {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}