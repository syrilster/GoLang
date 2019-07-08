package main

import "fmt"

type employee interface {
	display()
}

type FullTimeEmployee struct {
	name string
}

type ContractEmployee struct {
	name string
}

/*
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
Here the FullTimeEmployee type is implementing the employee interface to display the name
*/
func (emp FullTimeEmployee) display() {
	fmt.Println(emp.name)
}

func (emp ContractEmployee) display() {
	fmt.Print(emp.name)
}

func main() {
	var fullTimeEmp employee = FullTimeEmployee{"Tyson"}
	var contractEmp employee = ContractEmployee{"Mike"}
	fullTimeEmp.display()
	contractEmp.display()
}
