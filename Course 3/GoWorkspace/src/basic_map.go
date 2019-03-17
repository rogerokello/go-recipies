package main

import "./greeting"

func main() {

	//var s = greeting.Salutation{"1234567890", ""}

	slice := []greeting.Salutation{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}
	greeting.Greet(slice, greeting.MakePrinter("!!!!!"), true, 10)
}
