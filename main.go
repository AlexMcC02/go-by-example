package main

import (
    "fmt"
    "testing"
)

// The testing package provides the tools we need to write unit tests, with
// the `go test` command able to run them.

// As a rule of thumb, the test code will live alongside the code it is testing.

// This is the function we'll use for the purpose of demonstration.
func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// A function becomes a test when its signature starts with 'Test'.
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, - 2)
    if ans != -2 {
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// Writing tests can become repetitive, so it is idomatic to make use of a
// table-driven style, where test inputs and expected outputs are listed in
// a table and a single look walks over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%d, %d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

// Benchmark tests are used to ascertain the runtime of a piece of code, making
// them useful to identify performance bottlenbecks. The loop will be executed
// many times with an average returned to minimise moment-to-moment performance
// fluctuations.
func BenchmarkIntMin(b *testing.B) {
    for b.Loop() {
        IntMin(1, 2)
    }
}