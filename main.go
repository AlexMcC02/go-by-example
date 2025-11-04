package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct { // A singly linked-list, this time taking advantage of Go's iterators.
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val :v}
		lst.tail = lst.tail.next
	}
}

// Reminder: an iterator provides a way to access and traverse and access the elements of a collection, 
// one by one, without exposing the underlying data structure.

// All returns an iterator, which in Go is a function with a special signature.
// The iterator function takes another function as a parameter. By convention, this function is typically named yield.
// The iterator will then call yield for every element we want to iterate over. 
// You can take advantage of yield's return value for an early termination.
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// This function generates fibonacci numbers, and it will keep running as long as yield returns true.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	
	//List.All returns an interator, meaning it can be used in a regular range loop.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// The Slices package contains the Collect function, which will take any iterator and collects all its values into a slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// As soon as the loop hits break or early return, the yield function passed to the iterator will return false.
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}