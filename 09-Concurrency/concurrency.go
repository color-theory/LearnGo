/**
 * Concurrency
 * deadlock = one thread has locked a piece of data but needs to also see another piece of data that is locked by a different thread.
 * Go uses communication channel instead of data itself to control action. Don't share data. It's messy.
 * Do not communicate by sharing memory; instead, share memory by communicating.
 * Go Routines
 * Lightweight thread. Managed by go runtime. Uses Keyword "go"
 * channels can be buffered, which allows them to keep pushing messages until the buffer is full, then it is blocked
 * for range can be used to keep pulling from a channel, but if it is never closed, it will be a deadlock
 * Select is like a switch, but on communications.
 * it will execute a case that is "ready"
 * if more than one is ready, then execute random
 * if none are ready, block unless default is defined.
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

	c := make(chan greeting.Salutation) // creating a channel that returns Salutation items
	c2 := make(chan greeting.Salutation) // creating a second channel that returns Salutation items
	go salutations.ChannelGreeter(c) // start the goroutine
	go salutations.ChannelGreeter(c2) // start the second goroutine

	for{
		select{
			case s, ok := <- c:
				if ok{
					fmt.Println(s, ": 1")
				}else{
					return
				}
			case s, ok := <- c2:
				if ok{
					fmt.Println(s, ": 2")
				}else{
					return
				}
			default:
				fmt.Println("waiting...")
		}
	}
}
