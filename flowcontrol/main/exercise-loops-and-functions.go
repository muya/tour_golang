package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var l float64 = float64(1)
	var previousVal float64 = 0.0

	for z := 1; z <= 20; z++ {
		l -= (l*l - x) / (2 * l)
		fmt.Println(fmt.Printf("Iteration #%d | L: %f | PV: %f", z, l, previousVal))
		if previousVal == l || math.Abs(previousVal - l) < 0.01 {
			fmt.Println(fmt.Printf("Completing after %d iterations", z))
			break
		} else {
			previousVal = l
		}
	}

	return l
}

func main() {
	const ValD = 600000000
	fmt.Println(fmt.Sprintf("Our value: %f\n", Sqrt(ValD)))
	fmt.Println(fmt.Sprintf("Math value: %f\n", math.Sqrt(ValD)))
}
