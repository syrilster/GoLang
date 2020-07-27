package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
https://medium.com/swlh/go-common-misconceptions-about-goroutines-9dfa4bca3ba8
*/
type compute struct {
	sync.RWMutex
	response map[int]int
}

func fib(i int) int {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return fib(i-1) + fib(i-2)
}

func (c *compute) fibWorker(buffChan chan int, wg *sync.WaitGroup) {
	for i := range buffChan {
		result := fib(i)
		c.Lock()
		c.response[i] = result
		c.Unlock()
	}
	wg.Done()
}

func (c *compute) sequentialFibonacci(count int) {
	for i := 0; i < count; i++ {
		c.response[i] = fib(i)
	}
}

func (c *compute) parallelFibonacci(count int) {
	buffChan := make(chan int, 100)
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go c.fibWorker(buffChan, wg)
	}

	for i := 0; i < count; i++ {
		buffChan <- i
	}
	close(buffChan)
	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(4)
	action := flag.String("action", "s", "p for parallel and s for sequential")
	flag.Parse()
	startTime := time.Now()
	c := compute{
		response: make(map[int]int, 40),
	}
	switch *action {
	case "s":
		c.sequentialFibonacci(40)
	case "p":
		c.parallelFibonacci(40)

	}
	fmt.Println("Time Taken: ", time.Since(startTime).Milliseconds())
}
