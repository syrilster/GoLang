package main

import (
	"fmt"
	"time"
)

type CustomError struct {
	When time.Time
	What string
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("at %v, %s",
		err.When, err.What)
}

/*
	The error built-in interface type is the conventional interface for
	representing an error condition, with the nil value representing no error.
	type error interface {
		Error() string
	}
*/
func run() error {
	return &CustomError{
		time.Now(),
		"Sorry, It din't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
