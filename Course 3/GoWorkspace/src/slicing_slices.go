package main

import (
	"fmt"

	"./greeting"
)

func main() {

	s := []int{
		1, 2, 3, 4, 5,
	}

	// Print a slice of slice.
	// 1 denotes the starting position and 2 denotes the end but not inclusive
	fmt.Println(s[1:2])

	// leave out the end
	// This will slice from index 1 to the end
	fmt.Println(s[1:])

	// leave out the beginning
	// This will slice from the beginning to index 2(1 less than 3)
	fmt.Println(s[:3])

	slice := []greeting.Salutation{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}
	greeting.Greet(slice, greeting.MakePrinter("!!!!!"), true, 10)
}
