package main

import "fmt"

func main() {
	fmt.Println("Mutating maps")

	simpleMap := make(map[string]int)
	simpleMap["test"] = 100
	fmt.Println("The value in Map is: ", simpleMap)

	simpleMap["test"] = 42
	fmt.Println("The updated value in Map is: ", simpleMap["test"])

	delete(simpleMap, "test")
	fmt.Println("The value in Map is: ", simpleMap)

	//Test that a key is present with a two-value assignment:
	//If key is in simpleMap, isPresent is true. If not, isPresent is false.
	//If key is not in the simpleMap, then mapValue is the zero value for the map's element type.
	mapValue, isPresent := simpleMap["test"]
	fmt.Println("The value in Map ", mapValue, "is present? ", isPresent)
}
