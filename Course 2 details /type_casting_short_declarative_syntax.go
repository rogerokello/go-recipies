package main

import (
	"fmt"
	"reflect"
)

func main() {
	// This short form of assignment only works with in functions and only if
	// we are assigning values to variables at declaration
	a := 10.0000
	b := 3

	fmt.Println("A is of type ", reflect.TypeOf(a), " and B is of type ", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("C has value ", c, " and is of type ", reflect.TypeOf(c))

}
