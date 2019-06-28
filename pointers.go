package main

import "fmt"

/*
The & operator generates a pointer to its operand.
The * operator denotes the pointer's underlying value.
Unlike C, Go has no pointer arithmetic.
*/
func main() {
	randomNumber := 44
	fmt.Printf("The value of the number is %d \n", randomNumber)
	changeMeNow(&randomNumber)
	fmt.Printf("The value of the number is now %d", randomNumber)
}

func changeMeNow(input *int) {
	*input = 22
}
