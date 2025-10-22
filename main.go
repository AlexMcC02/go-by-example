package main

import (
	"fmt"
	"time"
)

func main() { 
	i := 2
	fmt.Print("Write ", i, " as " )
	switch i { // There is no need to add a break to the end of a switch case.
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // Commas can be used to separate multiple expressions in the same case.
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12: // Conditional case.
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch compares types instead of values. You can use this 
	// to discover the type of an interface value.
	whatAmI := func(i interface{}) { 
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}