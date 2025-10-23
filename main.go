package main

import "fmt"

// SlicesIndex is an example of a generic function, with it taking a slice of any comparable
// type and an element of that type and returns the index of the first ocurrence fo v in s, or
// -1 if not present. Note: being comparable means you can apply the == and != operators.
func SlicesIndex[S ~ []E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// List is an example of a generic type, it being a singly linked-list with values of any type.
type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val T
}

// We can define methods on generic types but we must keep the type parameters in place.
// E.g. the type is List[T] not List.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements returns all the List elements as a slice.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// When invoking generic functions, we can often rely on type inference.
func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex(s, "zoo")
	
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}