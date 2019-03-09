package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 1, 4)
	fmt.Printf("The length and capacity of mySlice is %d and %d respecively",
		len(mySlice), cap(mySlice))

	for i := 1; i <= 20; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nThe length and capacity of mySlice is %d and %d respecively",
			len(mySlice), cap(mySlice))
	}

	// Appending a slice to  Slice
	newSlice := []int{1, 3, 5, 7, 9}
	// use an elipsis(...) to extract the elements of the new slice
	mySlice = append(mySlice, newSlice...)
	fmt.Println("\n", mySlice)

}
