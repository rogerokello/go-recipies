package main

import (
	"fmt"
)

func main() {
	/* Declares a slice called myCourses
		make(<type>, <length>, <capacity>)
	  <capacity> : This defines the maximum size of the slice. Behind the scenes,
	  this is the size of the array that is going to back the slice. Every time
	  we create a slice go creates an array to store data for the slice. Slice
	  data is always stored in an array behind the scenes
	  <length> : This is the default length that will be stored with initial
	  values for the type. For example. In the case below, the first 5 values
	  in the newly created array will be given a default value of 0 since
	  this is an integer slice.
	*/
	myCourses := make([]int, 5, 10)
	fmt.Println("myCourses has a length of ", len(myCourses), " and capacity of ", cap(myCourses))
	for _, course := range myCourses {
		fmt.Println(course)
	}

	/*
	  This is a short form of creating a slice. What will happen in this case is that
	  the new slice will be created with a default size and length of the same size
	  for the underlying array supporting the slice
	*/

	myIntegers := []int{}
	fmt.Println("myIntegers has a length of ", len(myIntegers), " and capacity of ", cap(myIntegers))
	for _, value := range myIntegers {
		fmt.Println(value)
	}

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Slice indices are 0 based so the code below will print 5 and not 4
	fmt.Println("The value at index 4 is: ", mySlice[4])

	// This is changing the element at a particular position in a slice
	mySlice[2] = 20
	fmt.Println("The value at index 2 changed from  to: ", mySlice[2])

	/*Slice of a Slice
		- This will not include the value at position 5. it will
		  only include the values at index positions 1,2,3,4,
	  - if one does not specify slice start position, 0 is assumed eg [:4].
	  - if one does not specify slice end position, len(array) - 1 is assummed eg[4:]
	  - NB: Remember slices contain references to the original array, so changing
	    them will changed the original array and all slices referencing them will
	    see the changes
	*/
	sliceOfSlice := mySlice[1:5]
	fmt.Println("This is the first value of a slice of a slice: ", sliceOfSlice[0])
	fmt.Println("These are the values in the slice", sliceOfSlice)

	// Print this string in reverse order
	one := "This is me"
	arrayLen := len(one) - 1
	fmt.Print("The reverse of ", one, " is: ---->  ")
	for i := arrayLen; i >= 0; i-- {
		fmt.Print(string(one[i]))
		if i == 0 {
			fmt.Println("\n")
		}
	}
}
