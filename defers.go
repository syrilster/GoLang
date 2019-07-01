package main

import "fmt"

//A defer statement defers the execution of a function until the surrounding function returns.
// Like Java finally
func main() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}
