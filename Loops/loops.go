/**
 * Loops
 * Loops in go only use the keyword for but can simulate many of the common loop types
 */
package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Amy", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!~<3"), true, 5) //run greet 5 times
}
