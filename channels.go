package main

import "fmt"

func sum(input []int, channel chan int) {
	sum := 0
	for _, value := range input {
		sum += value
	}

	channel <- sum
}

/*
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
*/
func main() {
	randArray := []int{10, 20, 30, 4, 5, 6}
	channel := make(chan int)

	go sum(randArray[:len(randArray)/2], channel)
	go sum(randArray[len(randArray)/2:], channel)

	// This actually waits until the channel puts in some value.
	receiverOne, receiverTwo := <-channel, <-channel // receive from channel
	fmt.Println(receiverOne, receiverTwo, receiverOne+receiverTwo)
}
