package main

import (
	"fmt"
	"time"
)

var currentTime time.Time

func init() {
	currentTime = time.Now()
}

func main() {
	message := sayHello("Hello World :)")
	fmt.Println(message)
	fmt.Println("Current Time is --> ", currentTime.String())
}

func sayHello(message string) string {
	return message
}
