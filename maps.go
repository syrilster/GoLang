package main

import "fmt"

type Employee struct {
	age, salary int
}

var employeeMap map[string]Employee

var employeeMapLiteral = map[string]Employee{
	"Bell Labs": {40, 74},
	"Google":    {37, 122},
}

func main() {
	//The make function returns a map of the given type, initialized and ready for use.
	employeeMap = make(map[string]Employee)
	employeeMap["syril"] = Employee{30, 200}
	employeeMap["anju"] = Employee{28, 250}

	fmt.Println(employeeMap)
	fmt.Println(employeeMap["syril"])
	fmt.Println("Employee Literal Map: ", employeeMapLiteral)
}
