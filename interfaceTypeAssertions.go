package main

import "fmt"

func main() {
	var i interface{} = "Hello"

	/*
		A type assertion provides access to an interface value's underlying concrete value.

		t := i.(T)

		This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
		If i does not hold a T, the statement will trigger a panic.
	*/
	exampleOne := i.(string)
	fmt.Println(exampleOne)

	exampleTwo, Ok := i.(string)
	fmt.Println(exampleTwo, Ok)

	invalidExample, ok := i.(float64)
	fmt.Println(invalidExample, ok)

	//panic due to invalid type
	invalidExample = i.(float64)
	fmt.Println(invalidExample)
}
