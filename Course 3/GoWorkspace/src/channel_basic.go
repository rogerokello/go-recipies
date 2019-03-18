package main

import (
	"fmt"

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

	//Use fmt.Fprintf() writer interface
	// Synopsis
	// fmt.Fprintf(w io.Writer,a ...interface{}) (n int, err error)
	fmt.Fprintf(&salutations[0], "Brian number %d", 2)

	salutations = salutations.Add(greeting.Salutation{"James", "Oh hi"})

	done := make(chan bool)

	go func() {
		salutations.Greet(greeting.MakePrinter(" --Goroutine-- "), true, 10)
		done <- true
	}()

	salutations.Greet(greeting.MakePrinter("!!!!!"), true, 10)

	<-done
}
