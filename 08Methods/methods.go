/**
 * Methods
 * Methods have a type declared that it can act on
 * Go Methods can work on any named type(including built in types)
 * Go methods can operate on pointers to named types.
 */
package main

import "./greeting"

func main() {

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Mary", "What is up"},
	}

	salutations[0].Rename("John") // rename the Salutation at the 0 index

	salutations.Greet(greeting.CreatePrintFunction("!~<3"), true) //calling the Greet method on salutations.
}
