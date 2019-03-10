package main

import (
	"fmt"
)

func main() {

	// 'A' is different from "A" since 'A' is a rune type
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}

	// Running this for loop different times gives different results
	// because go randomises how it accesses the map. eg. Values in the
	// map will not be accessed in a particular order
	for key, value := range testMap {
		fmt.Printf("\nKey is %v and Value is %v", key, value)
	}
}
