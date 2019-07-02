package main

import "fmt"

type NilInterface interface {
	print()
}

type NilType struct {
	name string
}

func (elem *NilType) print() {
	if elem == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(elem.name)
}

/*
 In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully
handle being called with a nil receiver (as with the method print() in this example.)
*/
func main() {
	var nilInterface NilInterface
	var nilType *NilType

	nilInterface = nilType
	nilInterface.print()

	nilInterface = &NilType{"Hello"}
	nilInterface.print()
}
