package main

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	name     string
	greeting string
}

// Printer is a type which a function that takes in a string and returns nothing
type Printer func(string) // or type Printer func(string)()

// Greet takes in a function as a parameter
func Greet(salutation Salutation, do Printer) {
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
