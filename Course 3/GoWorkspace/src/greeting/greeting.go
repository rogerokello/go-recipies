package greeting

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	Name     string
	Greeting string
}

// Printer is a type which a function that takes in a string and returns nothing
type Printer func(string) // or type Printer func(string)()

// Greet takes in a function as a parameter
func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := "Mr "; isFormal {
		do(prefix + message)
	}
	do(alternate)
}

// CreateMessage returns named values
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY! " + name

	return
}

// MakePrinter is function that returns another function and retains its scope
func MakePrinter(custom string) Printer {
	return func(a string) {
		fmt.Println(a + custom)
	}
}
