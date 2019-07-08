package main

import (
	"fmt"
	"strings"
)

func main() {
	trimmedString, numberOfCharsRemoved := trimWithCount("This is a test string        ")
	fmt.Printf("%d characters were trimmed from [%s] \n", numberOfCharsRemoved, trimmedString)

}

func trimWithCount(inputString string) (string, int) {
	trimmedString := strings.Trim(inputString, " ")
	numberOfCharsRemoved := len(inputString) - len(trimmedString)
	return trimmedString, numberOfCharsRemoved
}
