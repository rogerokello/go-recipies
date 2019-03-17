package main

import "./greeting"

func main() {

	var s []int

	s = make([]int, 3, 4)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	slice := []greeting.Salutation{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}
	greeting.Greet(slice, greeting.MakePrinter("!!!!!"), true, 10)
}
