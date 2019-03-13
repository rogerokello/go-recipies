package main

import "fmt"

//Salutation  This is a definition of type string to be a Salutation
type Salutation string

func main() {

	var message Salutation = "Hello world"

	fmt.Println(message)
}
