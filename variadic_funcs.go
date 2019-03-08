package main

import (
	"fmt"
)

func main() {
	bestFinish := bestFinshes(34, 23, 12, 8, 11, 56, 56, 3, 3, 4)

	fmt.Println("Best league finish is ", bestFinish)
}

/*
This is a variadic function, it takes a variable number of arguments as specified
by the elipsis ..., finishes here is a slice of integers. The variable number of
arguments gets stored into the slice finishes
*/
func bestFinshes(finishes ...int) int {
	if len(finishes) > 0 {

		best := finishes[0]

		for _, item := range finishes {
			if item < best {
				best = item
			}
		}

		return best

	}

	return 0
}
