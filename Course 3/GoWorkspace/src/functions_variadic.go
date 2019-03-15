package main

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	name     string
	greeting string
}

// Greet is new type based on a struct
func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Yo")
	fmt.Println(message, alternate)
}

// CreateMessage is a variadic function which takes the first parameter name
// then after a variable number of arguments greeting. You must specify
// the type of the first parameter. In this case string
func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message = greeting[1] + " " + name
	alternate = "HEY! " + name

	return
}

func main() {
	var s = Salutation{"Roger", "Hello"}
	Greet(s)
}
