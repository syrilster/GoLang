package main

import "fmt"

type Employee struct {
	age, salary int
}

var employeeMap map[string]Employee

func main() {
	//The make function returns a map of the given type, initialized and ready for use.
	employeeMap = make(map[string]Employee)
	employeeMap["syril"] = Employee{30, 200}
	employeeMap["anju"] = Employee{28, 250}

	fmt.Println(employeeMap)
	fmt.Print(employeeMap["syril"])
}
