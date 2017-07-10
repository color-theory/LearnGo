package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool, times int) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)

	i := 0
	for i < times { // while i < times
		if prefix := GetPrefix(salutation.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
		i++  // increment i
	}
}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr "
		fallthrough
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
