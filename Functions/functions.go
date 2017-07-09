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
	_, alternate := CreateMessage(salutation.name, salutation.greeting) //_  ignores first return from CreateMessage, while alternate gets second returned string.

	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (string, string) { //create message takes name and greeting as string and returns two strings
	return greeting + " " + name, "Hey! " + name //returns first string followed by second string
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s) //calls Greet with s
}
