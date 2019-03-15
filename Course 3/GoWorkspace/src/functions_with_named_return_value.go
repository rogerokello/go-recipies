package main

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	name     string
	greeting string
}

// Greet is new type based on a struct
func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(message, alternate)
}

// CreateMessage returns named values
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY! " + name

	return
}

func main() {
	var s = Salutation{"Roger", "Hello"}
	Greet(s)
}
