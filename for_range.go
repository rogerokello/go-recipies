package main

import (
	"fmt"
)

func main() {

	// This creates and inialises a slice with default elements
	// the {} contains the elements to initialise it with
	coursesInProg := []string{"Docker", "Docker deep dive"}

	// This represents adding more elements to the end of the slice
	coursesInProg = append(coursesInProg, "Docker and Kubernetes")

	for _, value := range coursesInProg {
		fmt.Println(value)
	}
}
