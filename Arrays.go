package main

import "fmt"

func main() {
	//var fruitArr [2]string

	//Assigning values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Banana"

	//Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println("Array values ", fruitArr)
	fmt.Println("Second elem in array ", fruitArr[1])
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
