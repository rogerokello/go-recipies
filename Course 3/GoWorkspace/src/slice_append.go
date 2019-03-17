package main

import (
	"fmt"

	"./greeting"
)

func main() {

	s := []int{
		1, 2, 3, 4, 5,
	}

	// Append on item to the slice
	s = append(s, 1)
	fmt.Println(s)

	// Append all items in the previous slice using the ... operator
	// The ... basically says expand the slice
	// append actually takes a variable number of arguments of the type of slice
	s = append(s, s...)
	fmt.Println(s)

	slice := []greeting.Salutation{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}
	greeting.Greet(slice, greeting.MakePrinter("!!!!!"), true, 10)
}
