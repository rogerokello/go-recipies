package main

import (
	"fmt"
)

const (
	operatingSystem string = "MAC"
)

func main() {
	// This is a constant so you cannot reassign to it
	operatingSystem = 2

	fmt.Println("Meaning of the os is ", operatingSystem)
}
