package main

import (
	"fmt"
)

var maxSize int
var orderAmount int
var customerStatus string

func init() {
	maxSize = 6
	customerStatus = "ACTIVE"
	orderAmount = 5000
}

func main() {
	inputString := "Test"
	for index := 0; index < 10; index++ {
		fmt.Printf("Inside a go for loop take %d \n", index)
	}

	/*
		Go does NOT have the while or do loop control structures found in Java
		If you write a for statement with only a condition, and no initialization or post-iteration code,
		then it is essentially a Java while loop:
	*/
	for len(inputString) < maxSize {
		inputString = inputString + "Hello"
		fmt.Println("Inside a go while loop")
	}

	/*
		GO Switch:
		You can utilize multiple free-form conditions with each case clause.
		The conditions may be completely unrelated to each other, and can each be as complex as you like.
		Aside from flexible case conditions, the other key difference from Java’s switch is that Go’s version does NOT fall through.
	*/
	switch {
	case customerStatus == "ACTIVE":
		fmt.Println("Customer profile setup successfully")
	case orderAmount > 1000 && customerStatus == "Active":
		fmt.Println("Customer plan activated successfully")
	}
}
