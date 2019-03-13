package main

import "fmt"

func main() {

	message := "Hello world"
	var messagePtr = &message

	// This will modify what is being pointed to
	*messagePtr = "New Hello Message"

	fmt.Println(message, *messagePtr)
}
