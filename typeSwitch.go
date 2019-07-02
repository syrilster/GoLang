package main

import "fmt"

func do(value interface{}) {
	switch elem := value.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", elem, elem*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", elem, len(elem))
	default:
		fmt.Printf("I don't know about type %T!\n", elem)

	}
}

/*
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values),
and those values are compared against the type of the value held by the given interface value.
*/
func main() {
	do(100)
	do("Hello")
	do(true)
}
