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

}
