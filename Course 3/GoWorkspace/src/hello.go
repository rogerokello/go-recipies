package main

import "fmt"

func main() {

	message := "Hello world"
	var messagePtr = &message

	/*
	  - message is the message
	  - messagePtr is a pointer to the message
	  - *messagePtr is a way of dereferencing the message to get back the actual
	    message
	*/
	fmt.Println(message, messagePtr, *messagePtr)
}
