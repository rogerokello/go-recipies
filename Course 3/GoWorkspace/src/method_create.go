package main

import "./greeting"

func main() {

	salutations := greeting.Salutations{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}

	salutations.Greet(greeting.MakePrinter("!!!!!"), true, 10)
}
