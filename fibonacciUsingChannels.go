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

/*
Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the
receiver must be told there are no more values coming, such as to terminate a range loop.
*/
func main() {
	channel := make(chan int)
	go fibonacci(10, channel)
	for value := range channel {
		fmt.Println("Next in series is: ", value)
	}
}
