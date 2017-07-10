/**
 * Maps
 * Maps in go are like dictionaries. They have key value pairs
 * maps must be initialized with the make keyword before you can assign it values.
 */
package main

import "./greeting"

func main() {

	slice := []greeting.Salutation{ // setting up our slice
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Mary", "What is up"},
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("!~<3"), true) // call with our slice
}
