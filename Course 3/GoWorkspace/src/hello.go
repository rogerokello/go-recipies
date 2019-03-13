package main

import "fmt"

//Salutation  This is a definition of type string to be a Salutation
type Salutation struct {
	name     string
	greeting string
}

func main() {

	var s = Salutation{
		"Roger",
		"Hello ",
	}

	fmt.Println(s.name)
}
