package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

func (salutation *Salutation) Rename(newName string) { //must work on a pointer of Salutation or it will be working on a copy
	salutation.Name = newName
}

type Printer func(string) ()

type Salutations []Salutation //setting up our named type

func (salutations Salutations) Greet(do Printer, isFormal bool) { //changing our function to be a method that acts on Salutations type

	for i, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal {
			fmt.Print(i)
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr "

	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude"
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
