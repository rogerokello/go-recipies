package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(salutation.greeting)
	fmt.Println(salutation.name)
}

func main() {
	var s = Salutation{"Bob", "Hello"}

	Greet(s)

}
