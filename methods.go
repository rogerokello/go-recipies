package main

func main() {
	mp := messagePrinter{}
	mp.printMessage("Hi it Messy")
}

type messagePrinter struct {
	message string
}

/* The part (mp *messagePrinter) defines the context of the function
In this case it specifies that this is a method of the messagePrinter struct
*/
func (mp *messagePrinter) printMessage(message string) {
	println(message)
}
