package main

import "fmt"

func fibonacci(maxRange int, channel chan int) {
	x, y := 0, 1
	for index := 0; index < maxRange; index++ {
		channel <- x
		x, y = y, x+y
	}

	close(channel)
}

func main() {
	channel := make(chan int)
	go fibonacci(10, channel)
	for value := range channel {
		fmt.Println("Next in series is: ", value)
	}
}
