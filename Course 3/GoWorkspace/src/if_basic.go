package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Roger", "Hello"}
	greeting.Greet(s, greeting.MakePrinter("!!!!!"), true)
}
