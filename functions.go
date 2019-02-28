package main

func main() {
	message := "Hello"
	alteredMessage2 := "Before"

	sayHello(message)

	alterMessage(&alteredMessage2)
	println(alteredMessage2)

	variadicFunction("Hello", " my ", " name ", " is ")

	terms, result := add(1, 2, 4, 1, 8)
	println("The result of adding ", terms, "terms was", result)

	numTerms, sum := add2(1, 2, 4, 1)
	println("The result of adding ", numTerms, "terms was", sum)
}

func sayHello(message string) {
	println(message)
}

func alterMessage(message *string) {
	println(*message)
	*message = "After"
}

//Function taking in a variable number of args
func variadicFunction(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}

// Function with return values
func add(values ...int) (int, int) {
	result := 0
	for _, value := range values {
		result += value
	}

	return len(values), result
}

// Function with named return values
// more concise code
func add2(values ...int) (numTerms int, sum int) {
	for _, value := range values {
		sum += value
	}

	numTerms = len(values)

	return
}
