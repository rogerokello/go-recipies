package main

import (
	"fmt"
)

// Use the var key word if declaring at the package level
// Declaring variables like this makes them get
// initialised with a zero initial value
// 0 for a float64 and an empty string for a string type
var (
	name, course string
	module       float64
)

func main() {
	fmt.Println("The initail value of name is ", name)
	fmt.Println("The initial value for module is", module)
}
