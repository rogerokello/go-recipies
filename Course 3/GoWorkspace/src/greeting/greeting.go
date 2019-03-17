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
func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {
	// Use this form if you already know what is going to be in the map before hand
	prefixMap := map[string]string{
		"Roger": "Mr. ",
		"Joe":   "Dr. ",
		"Jenny": "Dr. ",
		"Mary":  "Mrs. ",
	}

	return prefixMap[name]
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
