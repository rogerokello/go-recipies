package main

func main() {
	mp := messagePrinter{"This is James"}
	mp.printMessage()

	//enmp := enhancedMessagePrinter{"This is James"} - This will not work
	enmp := enhancedMessagePrinter{}
	enmp.message = "This is Lenny"
	enmp.printMessage()

	//enmp.printMessage("Hi it is Lenny")
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

// Object composition in relation to
// inheritance
type enhancedMessagePrinter struct {
	messagePrinter
}
