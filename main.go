package main

import "fmt"

// This function passes by value, meaning it will get a copy of ival distinct from the one 
// in the calling function. As it does not return the value, the value passed in from the
// scope where the function is called will remain the same.
func zeroval(ival int) { 
	ival = 0
}

// Unlike the above function, this one specifies a *int parameter. This is an int pointer.
// The *iptr code will dereference the pointer from its memory address to the current value 
// at the address. Assigning a value to the dereferenced point changes the value at the
// referenced address.

// Simply put, this function, despite not returning anything, will be able to modify the
// value the pointer references.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// Zeroval does not change the i defined here, but zeroptr does as it has a reference to the
	// memory address for that variable.
	zeroval(i)
	fmt.Println("zeroval", i)

	zeroptr(&i) // The ampersand means you are providing the memory adress of i (or a pointer to i).
	fmt.Println("zeroptr", i)

	fmt.Println("pointer", &i)
}