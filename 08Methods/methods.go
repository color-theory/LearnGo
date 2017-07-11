/**
 * Methods
 */
package main

import "./greeting"

func main() {

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Mary", "What is up"},
	}

	slice = append(slice, greeting.Salutation{"Frank", "Hi"})
	greeting.Greet(slice, greeting.CreatePrintFunction("!~<3"), true)
}
