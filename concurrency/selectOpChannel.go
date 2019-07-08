package main

import "fmt"

/*
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/
func fibonacciUsingSelect(channel, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case channel <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func main() {
	normalChannel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-normalChannel)
		}
		quit <- 0
	}()

	fibonacciUsingSelect(normalChannel, quit)
}
