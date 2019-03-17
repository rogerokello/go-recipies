package main

import (
	"fmt"

	"./greeting"
)

func main() {

	s := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(s)

	slice := []greeting.Salutation{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}
	greeting.Greet(slice, greeting.MakePrinter("!!!!!"), true, 10)
}
