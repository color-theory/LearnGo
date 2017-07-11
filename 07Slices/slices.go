/**
 * Slices
 * Arrays in go are value types. The type is the size plus the underlying type. Arrays are fixed size, and zeroed out at init.
 * Slices are like an abstraction over an array.
 * size of slice doesn't affect its type. It's only based on underylying type, which allows us to pass different sized slices into a func
 * slice is basically a pointer to an array. It takes a slice of an array...looks at part of it.
 * slice is fixed size, but can be reallocated with append.
 * use make to init otherwise is nil
 * slices of slices possible. Changing the child slice will affect the parent slice.
 */
package main

import "./greeting"

func main() {

	var s []int
	s = make([]int, 3, 10) // length of 3 and capacity of 10 in the array...can expand length to capacity.
	s[0] = 1
	s[1] = 10
	s[2] = 500
	//s[3] = 23  out of range

	slice := []greeting.Salutation{ // setting up our slice
		{"Bob", "Hello"},
		{"Joe", "Hey"},
		{"Mary", "What is up"},
	}

	slice = append(slice, greeting.Salutation{"Frank", "Hi"}) // adding frank to slice
	slice = append(slice, slice...)                           //double up

	//	slice = slice[1:2] inclusive on start index, exclusive on ending index. Joe Only
	//	slice = slice[:2]  Bob and Joe
	slice = slice[1:]                                                 // Joe and Mary
	greeting.Greet(slice, greeting.CreatePrintFunction("!~<3"), true) // call with our slice
}
