package main

import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module       float64
	// These are inferred types. Go is clever enough to do this
	// Do not initialise variables in the in one line
	// since it makes readability difficult
	nameInferred, moduleInferred = "Infered string", 12.3
)

func main() {
	fmt.Println("Name is: ", name, "and type of Name is ", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and type of module is", reflect.TypeOf(module))
	fmt.Println("nameInferred is", nameInferred, "and type of nameInferred is", reflect.TypeOf(nameInferred))
	fmt.Println("moduleInferred is", moduleInferred, "and type of moduleInferred is", reflect.TypeOf(moduleInferred))
}
