package main

import "fmt"

func main() {
	/*
	  Go lets you define custom types using the type keyword. Any custom
	  type can have a set of methods associated with it.
	  Every field with in a single struct has to be unique
	*/

	// Here we have just defined a type. We have not yet initialised or
	// defined any course variables based in it
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// Declaring a new variable based on the struct we have defined.
	// This actually creates a struct with its values initialised with the
	// corresponding defaults for the types

	/*
		     // 1st way to do it - Using the var key word
			   var DockerDeepDive1 courseMeta
	*/

	/*
		    // 2nd way using the new keyword. Using the new keyword gives us a pointer
			  // to the newly created type
			    DockerDeepDive2 := new(courseMeta)
	*/

	// 3rd way is the composite literal way
	DockerDeepDive3 := courseMeta{
		Author: "Roger Okello",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println("DockerDeepDive3 value: ", DockerDeepDive3)
}
