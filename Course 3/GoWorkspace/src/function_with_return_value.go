package main

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	name     string
	greeting string
}

// Greet is new type based on a struct
func Greet(salutation Salutation) {
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

// CreateMessage is a way if creating a message on the fly
func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

func main() {
	var s = Salutation{"Roger", "Hello"}
	Greet(s)
}
