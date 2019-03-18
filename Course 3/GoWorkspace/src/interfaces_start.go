package main

import (
	"./greeting"
)

func renameToDog(renameInterface greeting.Renamable) {
	renameInterface.Rename("Dog")
}

func main() {

	salutations := greeting.Salutations{
		{"Roger", "Hi"},
		{"Joe", "Hello"},
		{"Mary", "Hi there"},
	}

	// The code below could be written as (&salutations[0]).Rename("Peter")
	salutations[0].Rename("Peter")

	// We nee to pass in an address to the slice because the Rename method of the
	// struct takes in a salutation pointer as its first parameter
	renameToDog(&salutations[0])

	salutations = salutations.Add(greeting.Salutation{"James", "Oh hi"})

	salutations.Greet(greeting.MakePrinter("!!!!!"), true, 10)
}
