package main

import "fmt"

const (
	// A incrementing iota's
	A = iota
	B
	C
)

func main() {
	fmt.Println(A, B, C)
}
