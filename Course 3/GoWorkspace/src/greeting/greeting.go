package greeting

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	Name     string
	Greeting string
}

// Printer is a type which a function that takes in a string and returns nothing
type Printer func(string) // or type Printer func(string)()

// Greet takes in a function as a parameter
func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	switch { // without name here it evaluates to true so comparisons are made on the inside for all values
	case name == "Roger": // All the cases have to evaluate to boolean
		prefix = "Mr. "
	case name == "Joe", name == "Jenny", len(name) == 10:
		// These are evaluated on an or basis.
		// Either name is joe or it Jenny or it's length is 10
		prefix = "Dr. "
	case name == "Mary":
		prefix = "Mrs. "
	default:
		prefix = "Dude "
	}

	return
}

// CreateMessage returns named values
func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY! " + name

	return
}

// MakePrinter is function that returns another function and retains its scope
func MakePrinter(custom string) Printer {
	return func(a string) {
		fmt.Println(a + custom)
	}
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
	}
}
