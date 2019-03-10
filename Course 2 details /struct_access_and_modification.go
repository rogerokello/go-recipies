package main

import "fmt"

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	DockerDeepDive := courseMeta{
		Author: "Roger Okello",
		Level:  "Intermediate",
		Rating: 5,
	}

	//Accesing individual fields in the struct is done using the dot(.) operator
	fmt.Println("Docker Deep dive author is ", DockerDeepDive.Author)

	//Individual fields can also be modified.
	DockerDeepDive.Rating = 6
	fmt.Println("Docker Deep dive has a new rating", DockerDeepDive.Rating)

}
