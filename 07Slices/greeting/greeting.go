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

	for i, s := range salutation {
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
