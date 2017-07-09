package main

import "fmt"

type salutation struct {
	name     string //structs are collections of types that are fields of structs. Can be used like a class, but classes do not exist in Go
	greeting string
}

const (

	PI       float32 = 3.14  //iota 0
	Language         = "Go" //the type of constants do not need to be specified during assignment. iota 1
	A                = iota //iota starts at 0 and increments by 1 each item in const. It resets to 0 whenever const() is called. iota 2

)

const (
	B = iota //B should be 0 again because it is in a new const
	C        // and C will be assigned the same value as above

	// causing the A B and C to be 2 0 and 1
)

func main() {

	var message = salutation{"Saki", "hello!"}
	var toPrint *salutation = &message //toPrint is a pointer of type salutation that holds the memory address of message

	*toPrint = salutation{"沙記", "こんにちは、"} //the value at the memory address that toPrint holds is being set to a new salutation

	fmt.Println(toPrint, &toPrint, &message, &toPrint.name, &message.name) // go gives you shorthand for using pointers of structs.

	fmt.Println(message) //message should have changed when the value was changed using the toPrint pointer above.
	fmt.Println(message.greeting, message.name)

	fmt.Println(PI)
	fmt.Println(Language)

	d := 42                  // This is the same as var d int = 42
	var e, f int = 3, 4      // can assign multiple vars are once
	var g, h = 77, "goodbye" // multiple assignments to different types is possible

	fmt.Println(A, B, C, d, e, f, g, h) //vars are case sensitive.
}
