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

	ch1 := make(chan greeting.Salutation)
	ch2 := make(chan greeting.Salutation)

	go salutations.ChannelGreeter(ch1)

	go salutations.ChannelGreeter(ch2)

	for {
		select {
		case s, ok := <-ch1:
			if ok {
				fmt.Println(s)
			}
		case s, ok := <-ch2:
			if ok {
				fmt.Println(s)
			}
		default:
			fmt.Println("Waiting......")
		}
	}

}
