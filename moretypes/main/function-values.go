package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypt := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("Hypotenuse of 5, 12 is: %f\n", hypt(5, 12))

	fmt.Printf("Default hypot is: %f\n", compute(hypt))
	fmt.Printf("Hypot for math.Pow is: %f\n", compute(math.Pow))
}
