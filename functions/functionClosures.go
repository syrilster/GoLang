package main

import "fmt"

/*
	This function returns another function, which we define anonymously in the body of intSeq.
	The returned function closes over the variable number to form a closure.
*/
func intSeq() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}

func main() {
	/*
		We call intSeq, assigning the result (a function) to nextInt.
		This function value captures its own number value, which will be updated each time we call nextInt.
	*/
	nextNumber := intSeq()

	fmt.Println("The next number is: ", nextNumber())
	fmt.Println("The next number is: ", nextNumber())
	fmt.Println("The next number is: ", nextNumber())

	/*
		To confirm that the state is unique to that particular function, create and test a new one.
	*/
	newNextNumber := intSeq()

	fmt.Println("The next new number is: ", newNextNumber())
	fmt.Println("The next new number is: ", newNextNumber())
}
