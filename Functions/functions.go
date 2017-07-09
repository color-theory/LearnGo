/**
 * Multiple return values
 * Use it like any other type: Similar to functions in js
 * Function literals: When declaring a function inside another function, it remembers context.
 */
package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	first, alternate := CreateMessage(salutation.name, salutation.greeting) //first takes first return from CreateMessage, while alternate gets second returned string.
	fmt.Println(first)
	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (message string, alternate string) { //CreateMessage takes name and greeting as strings and returns two named strings
	message = greeting + " " + name
	alternate = "Hey! " + name
	return //returns both message and alternate
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s) //calls Greet with s
}
