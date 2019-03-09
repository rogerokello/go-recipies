package main

import "fmt"

func main() {
	/*
			  map construct
		    NB: A map is an unordered list unlike a slice
			  map[keyType]valueType
			  Since go is strongly typed we must declare both the valueType and the keyType
			  when declaring a map.
			  NB.
			  1. The keyType must be a comparable type ie a type that can be compared with
			  the equal to(==) or (!= )
			  2. All keys in a map have to be unique. You cannot have two items with The
			  same key
	*/

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	//Creating 0f maps using the composite literal form
	recentHeadToHead := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\n League titles: %v \n Recent head to heads %v \n",
		leagueTitles, recentHeadToHead)
}
