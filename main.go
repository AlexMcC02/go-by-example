package main

import "fmt"

// An enumerated type has a fixed number of possible values, each with distinct names.
// Go does not have a distinct enum type, but they can be implemented using existing language idioms.

// The ServerState type is an int under the hood.
type ServerState int

// The possible values are defined as constants. The keyword iota generates successive constant values
// automatically -- 0, 1, 2, 3.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Here we implement the fmt.Stringer interface, allowing us to output values of ServerState as strings.
var stateName = map[ServerState]string {
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

// Though ServerState is implemented as an int, passing an int to transition will result in a type
// mismatch, providing some degree of compile-time type safety for enums.
func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// This function emulates a state transition for a web server, it takes the existing state 
// and returns a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unkwown state: %s", s))
	}
}