/**
 * Branching
 * if statements not much different from other languages
 * Does have optional statement followed by semicolon
 * no parentheses
 * * * *
 * switch statements are quite different
 * no default fall through
 * can specify a fall through, but it will break by default
 * expression not necessary - will evaluate true
 * cases can be expressions
 * can switch on types instead of value
 */
package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!closure!"), true)

	var t = greeting.Salutation{"Amy", "Hello"}
	greeting.Greet(t, greeting.CreatePrintFunction("!~<3"), true)

}
