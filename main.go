package main

import "fmt"

// Initialising a person struct with name and age fields.
type person struct {
	name string
	age int
}

// This function constructs a new person with a given name.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// Because go uses garbage collection, it is safe to return a pointer to a local variable.
	// It will only be cleaned up by the gc when there are no active references to it.
	return &p 
}

func main() {
	fmt.Println(person{"Bob", 20}) // Creating a new struct in-line.

	fmt.Println(person{name: "Alice", age: 30}) // You can specify the fields directly.

	fmt.Println(person{name: "Fred"}) // Omitted fields are zero-valued.

	// Remember that an ampersand yields a pointer, in this case to the struct.
	fmt.Println(&person{name: "Ann", age: 40}) 

	// It is considered idiomatic to encapsulate new struct creation in constructor functions
	// just like those used in more traditional OOP languages.
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Dot syntax used to access fields.

	sp := &s
	fmt.Println(sp.age) // Dots can be used with struct pointers, the pointers are dereferenced automatically.

	sp.age = 51 // Structs are mutable.
	fmt.Println(sp.age)

	dog := struct { // Anonymous struct type. This is useful for table-driven tests.
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}