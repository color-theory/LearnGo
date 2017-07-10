/**
 * Loops
 * Loops in go only use the keyword for but can simulate many of the common loop types
 */
package main

import "./greeting"

func main() {

	slice := []greeting.Salutation{  // setting up our slice
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Mary", "What is up"},
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("!~<3"), true) // call with our slice
}
