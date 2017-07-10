package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal { //an optional statement gets prefix based on name if isFormal.
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	switch { //this is now switching on a boolean equal to true
	case name == "Bob": // now needing to change case to an expression that evaluates to a bool
		prefix = "Mr "
		fallthrough // after setting prefix to "Mr ", it falls through and sets prefix to "Dr "
	case name == "Joe":
		prefix = "Dr "
	case name == "Mary", name == "Amy", len(name) == 10:
		prefix = "Mrs "
	default:
		prefix = "Dude "
	}
	return
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey! " + name
	return
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
