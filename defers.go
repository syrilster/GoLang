package main

import "fmt"

//A defer statement defers the execution of a function until the surrounding function returns.
// Like Java finally
func main() {

	fmt.Println("Hello")

	//Deferred function calls are pushed onto a stack. When a function returns,
	// its deferred calls are executed in last-in-first-out order.

	for index := 0; index < 10; index++ {
		defer fmt.Println(index)
	}

	fmt.Println("World")
}
