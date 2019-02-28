package main

import "fmt"

func main() {
	foo := newMyStruct()
	foo.myMap["one"] = "one"

	fmt.Println(foo.myMap["one"])
}

type myStruct struct {
	myMap map[string]string
}

func newMyStruct() *myStruct {
	result := myStruct{}
	result.myMap = make(map[string]string)

	return &result
}
