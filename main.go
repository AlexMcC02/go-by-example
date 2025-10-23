package main

import "fmt"


// Go supports embedding of structs and interfaces to express a seamless composition of types.
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// The container embeds a base. 
type container struct {
	base // An embedding looks like a field without a name.
	str string
}

// When creating structs with literals, we need to initialise the embedding explicitly, here the
// embedded type servers as the field name.
func main() {
	co := container {
		base: base {
			num: 1,
		},
		str: "some name",
	}

	// We can access the base's fields directly on co, without needing to specify its parent.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Or, we can spell out the full path using the embedded type's name.
	fmt.Println("also num:", co.base.num)

	// Since contianer embeds base, the methods of base also become methods of a container.
	fmt.Println("describe", co.describe())

	// Embedding structs with methods may be used to bestow interface implementations onto other
	// structs. Here we see that a container now implements the describer interface because it itself
	// embeds base.
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}