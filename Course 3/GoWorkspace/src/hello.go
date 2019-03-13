package main

import "fmt"

//Salutation  This is a definition of type string to be a Salutation
// Capitalizing the S in salutation means that salutation will
// be visible outside this package
type Salutation struct {
	name     string
	greeting string
}

func main() {

	var s = Salutation{}

	s.name = "Roger"
	s.greeting = "Hello"

	fmt.Println(s.greeting, " ", s.name)
}
