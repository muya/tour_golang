package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// custom method
func (f MyFloat) Subd(v float64) float64 {
	return float64(f) - v
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Println(f.Subd(8))
}
