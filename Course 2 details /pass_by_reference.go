package main

import (
	"fmt"
)

func main() {
	name := "Roger"
	course := "Mendels things"

	fmt.Println(
		"Hi!  ",
		name,
		"You are watching ",
		course)

	/*
	   This will pass a value by reference
	   ie. it will create a copy of course and send it to the function. So the
	   function will be working with a copy and not the actual value
	*/
	changeCourse(&course)
	fmt.Println("\n You are now watching course ", course)
}

/*
The * means that variable course is a pointer to a string
In short the function takes in a pointer to a string
*/
func changeCourse(course *string) string {
	// This means we are dereferencing the pointer
	// making the change to what has been referenced
	*course = "First look, Native docker clustering"
	fmt.Println("Tyring to change course to ", *course)
	return *course
}
