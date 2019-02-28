package main

import "fmt"

func main() {
	addition := 1 + 2

	subtraction := 2 - 1

	division := 2 / 1

	multiply := 2 * 1

	remainder := 5 % 2

	fmt.Println(
		"Add", addition,
		"Sub", subtraction,
		"division", division,
		"multiply", multiply,
		"remainder", remainder)

	inc := 1
	inc++
	fmt.Println("Increment", inc)

	// Augemented assignment(+) operator
	inc += 5
	fmt.Println("Augmented Assignment: --->", inc)

	// Augmented assignment(*) operator
	inc *= 5
	fmt.Println("Augmented Multiplication: -->", inc)

}
