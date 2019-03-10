package main

import (
	"fmt"
	"reflect"
)

/*
& references a pointer to something. ie gets a pointer to something
* dereferences a pointer to something. ie when placed before something that is
a pointer, it will get the contents of that thing.
*/
func main() {
	module := 3.2
	ptr := &module

	fmt.Println("Module is ", module, " and is of type ", reflect.TypeOf(module))
	fmt.Println("Module address is ", ptr, " and its referenced value through a pointer is ", *ptr)
}
