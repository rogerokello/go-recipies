package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.0000
	b := 3

	fmt.Println("A is of type ", reflect.TypeOf(a), " and B is of type ", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("C has value ", c, " and is of type ", reflect.TypeOf(c))

}
