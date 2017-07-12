package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type RenameType struct {
	Name string
}

type Renamable interface {
	Rename(newName string)
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

func (renametype *RenameType) Rename(newName string) {
	renametype.Name = newName
}

func (salutations Salutations) ChannelGreeter(c chan Salutation) {
	for _, s := range salutations{
		c <- s
	}
	close(c) // we must close this channel after we are done sending messages or else we will deadlock with all channels asleep.
}

type Printer func(string) ()

type Salutations []Salutation

func (salutations Salutations) Greet(do Printer, isFormal bool) {

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
