/**
 * Multiple return values
 * Use it like any other type: Similar to functions in js
 * Function literals: When declaring a function inside another function, it remembers context.
 * Variadic functions: multiple variable parameters in func
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
	first, alternate := CreateMessage(salutation.name, salutation.greeting, "yo!", "sup!") 	//first takes first return from CreateMessage,
																					// while alternate gets second returned string. Passing in a 3rd parameter as second greeting
	fmt.Println(first)
	fmt.Println(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) { //CreateMessage takes name and greeting as strings and returns two named strings

	message = greeting[len(greeting)-1] + " " + name //using last greeting passed "sup!"
	alternate = "Hey! " + name
	return //returns both message and alternate
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s) //calls Greet with s
}
