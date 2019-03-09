package main

import (
	"fmt"
)

func main() {
	// An if can have scoped variables. In this case the variables will
	// be only available with in the if block

	if one, two := 1, 2; one < two {
		fmt.Println("One is less than two")
	} else if one > two {
		fmt.Println("One is greater than two")
	} else if one == two {
		fmt.Println("One is equal to two")
	} else {
		fmt.Println("Some thing unusual happened")
	}
}
