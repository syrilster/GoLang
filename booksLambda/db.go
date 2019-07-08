package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-2"))

func getItem(isbn string) (*Book, error) {
	fmt.Println("Inside lambda function")
	fmt.Println("DB value is: ", db)
	input := &dynamodb.GetItemInput{
		TableName: aws.String("books"),
		Key: map[string]*dynamodb.AttributeValue{
			"ISBN": {
				S: aws.String(isbn),
			},
		},
	}

	fmt.Println("Fetching Record: ", input)

	result, err := db.GetItem(input)

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type map[string]*AttributeValue. We can use the UnmarshalMap
	// helper to parse this straight into the fields of a struct. Note:
	// UnmarshalListOfMaps also exists if you are working with multiple items.

	book := new(Book)
	err = dynamodbattribute.UnmarshalMap(result.Item, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// Add a book record to DynamoDB.
func putItem(book *Book) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Books"),
		Item: map[string]*dynamodb.AttributeValue{
			"ISBN": {
				S: aws.String(book.ISBN),
			},
			"Title": {
				S: aws.String(book.Title),
			},
			"Author": {
				S: aws.String(book.Author),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
