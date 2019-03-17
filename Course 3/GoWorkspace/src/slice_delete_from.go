package main

import (
	"fmt"

	"./greeting"
)

func main() {

	s := []int{
		1, 2, 3, 4, 5,
	}

	// This will delete everythin from index 2 to index 3.
	// ie from 2 to 3(4-1)
	s = append(s[:2], s[4:]...)
	fmt.Println(s)

	slice := []greeting.Salutation{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}
	greeting.Greet(slice, greeting.MakePrinter("!!!!!"), true, 10)
}
