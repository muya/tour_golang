package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Method call sin pointer: %f\n", v.Abs())
	fmt.Printf("fn call sin pointer: %f\n", AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Printf("Method call with pointer: %f\n", p.Abs())
	fmt.Printf("fn call with pointer: %f\n", AbsFunc(*p))
}
