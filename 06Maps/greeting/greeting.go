package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation []Salutation, do Printer, isFormal bool) {

	for i, s := range salutation { // loop through each item in salutation
		message, alternate := CreateMessage(s.Name, s.Greeting) // need to use s now

		if prefix := GetPrefix(s.Name); isFormal {
			fmt.Print(i) // print the index
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	var prefixMap map[string]string     //create variable for map
	prefixMap = make(map[string]string) //initialize map

	prefixMap["Bob"] = "Mr "
	prefixMap["Joe"] = "Dr "
	prefixMap["Amy"] = "Dr " // although keys must be unique, values can be duplicated
	prefixMap["Mary"] = "Mrs "

	return prefixMap[name]
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
