package main

import "fmt"

func main() {
	myString := []string{}
	fmt.Println(myString)

	myString = append(myString, "one")
	fmt.Println(myString[0])

	mySlice := make([]string, 4, 100)
	mySlice[0] = "12."
	mySlice[1] = "13."
	mySlice[2] = "14."
	fmt.Println(mySlice)

	myMap := make(map[int]string)

	myMap[0] = "Zero"
	myMap[1] = "One"
	myMap[3] = "Three"

	fmt.Print(myMap[0], "  " ,myMap[3])
}
