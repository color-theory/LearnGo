/**
 * Concurrency
 * deadlock = one thread has locked a piece of data but needs to also see another piece of data that is locked by a different thread.
 * Go uses communication channel instead of data itself to control action. Don't share data. It's messy.
 * Do not communicate by sharing memory; instead, share memory by communicating.
 * Go Routines
 * Lightweight thread. Managed by go runtime. Uses Keyword "go"
 */
package main

import (
	"./greeting"
	"fmt"
)

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Mary", "What is up"},
	}

	salutations[0].Rename("John")

	RenameToFrog(&salutations[0])

	renametype := greeting.RenameType{"myname"}
	fmt.Println(renametype)

	RenameToFrog(&renametype)
	fmt.Println(renametype)

	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	done := make(chan bool)

	go func() { // runs this as a go routine
		salutations.Greet(greeting.CreatePrintFunction("<CONCURRENT>"), true)
		done <- true // putting true onto the done channel
	}()

	salutations.Greet(greeting.CreatePrintFunction("!~<3"), true)

	<-done // block until getting data on the done channel. We could send that to a variable, but instead we just ignore the value
}
