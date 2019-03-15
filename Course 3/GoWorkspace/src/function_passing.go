package main

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	name     string
	greeting string
}

// Greet is new type based on a struct
func Greet(salutation Salutation, do func(string)) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	do(message)
	do(alternate)
}

// CreateMessage returns named values
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY! " + name

	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func main() {
	var s = Salutation{"Roger", "Hello"}
	Greet(s, PrintLine)
}
