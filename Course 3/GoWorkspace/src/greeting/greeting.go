package greeting

import "fmt"

// Salutation is new type based on a struct
type Salutation struct {
	Name     string
	Greeting string
}

type Renamable interface {
	Rename(newName string)
}

// This is a pointer reciever method and will modify the underlying value
// passed to it.
func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

// Printer is a type which a function that takes in a string and returns nothing
type Printer func(string) // or type Printer func(string)()

type Salutations []Salutation

// This is a method on the type Salutations
func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {
	for _, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func (salutations Salutations) ChannelGreeter(c chan Salutation) {
	for _, value := range salutations {
		c <- value
	}
	close(c)
}

func (salutations Salutations) Add(salutation Salutation) Salutations {
	s := append(salutations, salutation)
	return s
}

func GetPrefix(name string) (prefix string) {
	// Use this form if you already know what is going to be in the map before hand
	prefixMap := map[string]string{
		"Roger": "Mr. ",
		"Joe":   "Dr. ",
		"Jenny": "Dr. ",
		"Mary":  "Mrs. ",
	}

	prefixMap["Joe"] = "Jr. "

	// Delete syntax is delete(collection, key)
	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
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
