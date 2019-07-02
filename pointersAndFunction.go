package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (vertex Vertex) square() float64 {
	return math.Sqrt(vertex.x * vertex.y)
}

/*
This is a method of Vertex as opposed to a function scale. Check invocation below.
*/
func (vertex *Vertex) scale(factor float64) {
	vertex.x = vertex.x * factor
	vertex.y = vertex.y * factor
}

/*
This is a function scale. Check invocation below.
*/
func scale(vertex *Vertex, factor float64) {
	vertex.x = vertex.x * factor
	vertex.y = vertex.y * factor
}

func main() {
	vertex := Vertex{4, 4}
	fmt.Println("Invoking the method scale on vertex type:")
	/*
		For the statement vertex.scale(10), even though vertex is a value and not a pointer, the method with the pointer receiver is called automatically.
		That is, as a convenience, Go interprets the statement vertex.scale(10) as (&vertex).scale(10) since the scale method has a pointer receiver.
	*/
	vertex.scale(10)
	fmt.Println("Calling the function scale:")
	scale(&vertex, 10)
	fmt.Println("The Final result ")
	fmt.Println(vertex.square())
}
