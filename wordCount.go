package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)
	stringArray := strings.Split(s, " ")
	for index := 0; index < len(stringArray); index++ {
		elemValue, isPresent := wordCountMap[stringArray[index]]
		if isPresent {
			wordCountMap[stringArray[index]] = elemValue + 1
		} else {
			wordCountMap[stringArray[index]] = 1
		}
	}
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
