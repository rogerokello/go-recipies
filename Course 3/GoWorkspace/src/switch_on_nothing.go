package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"1234567890", ""}
	greeting.Greet(s, greeting.MakePrinter("!!!!!"), true)
}
