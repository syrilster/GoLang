package main

import (
	"fmt"
	"sync"
)

/*
This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.
Go's standard library provides mutual exclusion with sync.Mutex and its two methods:

	Lock
	Unlock
We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.
We can also use defer to ensure the mutex will be unlocked as in the Value method.
*/
type AtomicCounter struct {
	value map[string]int
	mutex sync.Mutex
}

func (counter *AtomicCounter) incrementValue(key string) {
	counter.mutex.Lock()
	counter.value[key]++
	defer counter.mutex.Unlock()
}

func (counter *AtomicCounter) getValue(key string) int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.value[key]
}

func main() {
	key := "some-key"
	counter := AtomicCounter{value: map[string]int{}}
	for index := 0; index < 1000; index++ {
		counter.incrementValue(key)
	}

	fmt.Println(counter.getValue(key))
}
