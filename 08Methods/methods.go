/**
 * Methods
 * Methods have a type declared that it can act on
 * Go Methods can work on any named type(including built in types)
 * Go methods can operate on pointers to named types.
 * Interfaces
 * No way to require that an interface be implemented.
 * As long as a type implements the methods defined in an interface, it implements that interface
 * Allows you to pass multiple types into a function by passing in an interface
 * an empty interface is implemented by any type.
 */
package main

import (
	"./greeting"
	"fmt"
)

func RenameToFrog(r greeting.Renamable) { // a function that works on a Renamable interface implementing type
	r.Rename("Frog")
}

func main() {

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Mary", "What is up"},
	}

	salutations[0].Rename("John") // rename the Salutation at the 0 index

	RenameToFrog(&salutations[0]) //must dereference salutations[0] as Rename works on a pointer to a Salutation

	renametype := greeting.RenameType{"myname"} //creating a RenameType
	fmt.Println(renametype)

	RenameToFrog(&renametype) // This workds because it is also a Renamable
	fmt.Println(renametype)

	fmt.Fprintf(&salutations[0], "The count is %d", 10) // this can work on salutations because it can be passed as a Writer now

	salutations.Greet(greeting.CreatePrintFunction("!~<3"), true) //calling the Greet method on salutations.
}
