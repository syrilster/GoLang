package main

import (
	"fmt"
	"github.com/pkg/errors"
)

//errors.New constructs a basic error value with the given error message.
func functionOne(argOne int) (int, error) {
	if argOne > 50 {
		return -1, errors.New("Can't process this number")
	}

	return argOne + 10, nil
}

type httpError struct {
	code    int
	message string
}

func (err *httpError) Error() string {
	return fmt.Sprintf("%d - %s", err.code, err.message)
}

func functionTwo(arg int) (int, error) {
	if arg == 42 {
		return -1, &httpError{500, "Internal server error"}
	}

	return arg + 10, nil
}

func main() {
	for _, index := range []int{1, 110} {
		if value, err := functionOne(index); err != nil {
			fmt.Println("function One did not work: ", err)
		} else {
			fmt.Println("function One worked and value is: ", value)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := functionTwo(i); e != nil {
			fmt.Println("function Two failed:", e)
		} else {
			fmt.Println("function Two worked:", r)
		}
	}
}
