package main

import (
	"fmt"
	"time"
)

func printString(input string) {
	for index := 0; index < 5; index++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(input)
	}
}

func main() {
	//Evaluation of go printString("Hello") happens in the current goroutine and the execution of printString("World) happens in the new goroutine.
	go printString("Hello")
	printString("World")
}
