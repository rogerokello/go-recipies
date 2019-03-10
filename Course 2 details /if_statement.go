package main

import (
	"fmt"
)

func main() {
	variableOne := "True"

	variableTwo := "Flask"

	// The condition in the if must evaluate to a boolean true or false
	// It cannot be like an a string or something like a datastructure with
	// things in it.
	// You could also use angle brackets if conditions are many
	if variableOne == "t" && (variableTwo == "Flask" || true) && true {
		fmt.Println("Variable one is true")
	} else if variableOne != "Flask" {
		fmt.Println("Variable one is not true")
	} else {
		fmt.Println("Other situation")
	}
}
