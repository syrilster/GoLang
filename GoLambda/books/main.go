package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Book struct {
	ISBN   string `json:isbn`
	Title  string `json:title`
	Author string `json:author`
}

func show() (*Book, error) {
	firstBook := &Book{
		ISBN:   "97833-1420931693",
		Title:  "The Alchemist",
		Author: "Paulo Coelho",
	}

	return firstBook, nil
}

func main() {
	lambda.Start(show)
}
