package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Here the switch statement is using an initializer value
	// to get a random value which will be used during the comparison operation
	// A case statement may also choose to match multiple values
	// as is specifiied below in the second case which chooses to match any of
	// the strings "odd", "prime"
	value := randomGenerator()
	switch valueType := checkType(value); valueType {
	case "even":
		fmt.Println("The number ", value, ", generated is even")
	case "odd", "prime":
		fmt.Println("The number ", value, ", generated is ", valueType)
	}

}

func randomGenerator() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}

func checkType(value int) string {
	if value%2 == 0 {
		return "even"
	}

	return "odd"
}
