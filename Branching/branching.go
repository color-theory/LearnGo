/**
 * Branching
 * if statements not much different from other languages
 * Does have optional statement followed by semicolon
 * no parentheses
 */
package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!closure!"), false)
}
