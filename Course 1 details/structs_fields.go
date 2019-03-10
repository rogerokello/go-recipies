package main

import "fmt"

func main() {
	// Initialise a type with values
	// First parameter represents the first field in the
	// type. In this case the myField
	// This method creates an object in the local execution
	// stack
	foo1 := myStruct{"Value for field"}
	println(foo1.myField)

	/* Initialise with defaults for each of the values
		   of the struct
	     This also creates an object in the local
	     execution stack
	*/
	foo2 := myStruct{}
	foo2.myField = "Value for the field of foo2"
	println(foo2.myField)

	/*
	   use the new keyword such that values are put on the heap
	   instead of the call stack.
	   It is often better to create large objects on the heap
	   instead
	*/
	foo3 := new(myStruct)
	// In this case foo3 is a memory address and referencing its
	// attributes through object notation dereferences the pointer
	foo3.myField = "Value for field of foo3"
	println(foo3.myField)

	foo4 := new(myStruct)
	println(foo4)     // This is actually a pointer it prints an address
	fmt.Println(foo4) // This prints a refence to an empty object

	foo5 := &myStruct{"foo5"}
	println(foo5)
	fmt.Println(foo5)

	foo6 := myStruct{}
	foo6.myField = "foo6"
	// println(foo6) - This wont work
	fmt.Println(foo6)
}

type myStruct struct {
	myField string
}
