package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
	"regexp"
)

var isbnRegex = regexp.MustCompile(`[0-9]{3}\-[0-9]]{10}`)
var errorLogger = log.New(os.Stderr, "ERROR", log.Llongfile)

type Book struct {
	ISBN   string `json:isbn`
	Title  string `json:title`
	Author string `json:author`
}

func router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return show(request)
	case "POST":
		return create(request)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusMethodNotAllowed)
	}

	book := new(Book)
	err := json.Unmarshal([]byte(request.Body), book)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if !isbnRegex.MatchString(book.ISBN) {
		return clientError(http.StatusBadRequest)
	}

	if book.Title == "" || book.Author == "" {
		return clientError(http.StatusBadRequest)
	}

	err = putItem(book)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/books?isbn=%s", book.ISBN)},
	}, nil

}

func show(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Get the ISBN from the query parameters
	isbn := request.QueryStringParameters["isbn"]
	if !isbnRegex.MatchString(isbn) {
		return clientError(http.StatusBadRequest)
	}
	book, err := getItem(isbn)
	if err != nil {
		return serverError(err)
	}

	if book == nil {
		return clientError(http.StatusNotFound)
	}
	bookJson, err := json.Marshal(book)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bookJson),
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}
