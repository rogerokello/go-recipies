package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Jenny", "Hello"}
	greeting.Greet(s, greeting.MakePrinter("!!!!!"), true)
}
