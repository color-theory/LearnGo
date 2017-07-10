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
	for { // forever
		if i >= times{
			break //break out of the loop
		}

		if i % 2 == 0 { // if there is no remainder when dividing by 2. Even number
			i++
			continue // jump back up to the beginning of the loop
		}
		if prefix := GetPrefix(salutation.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
		i++
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
