package main

import (
	"fmt"
)

func main() {
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	// Accessing a map value
	fmt.Println("Item at map index A is ", testMap["A"])

	// Printing the whole map
	fmt.Println("Map has values: ", testMap)

	// Modifying an Existing value
	testMap["B"] = 100
	// Adding a non existing value
	testMap["F"] = 12345
	fmt.Println("New values for the map: ", testMap)

	// Deleting an element from the map. Use the in built delete method
	// Format of delete is always delete(<map>, <key>)
	delete(testMap, "C")
	fmt.Println("Currently existing look of map: ", testMap)

}
