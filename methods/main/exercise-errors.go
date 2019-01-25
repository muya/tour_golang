package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var l float64 = float64(1)
	var previousVal float64 = 0.0

	for z := 1; z <= 20; z++ {
		l -= (l*l - x) / (2 * l)
		fmt.Println(fmt.Printf("Iteration #%d | L: %f | PV: %f", z, l, previousVal))
		if previousVal == l || math.Abs(previousVal-l) < 0.01 {
			fmt.Println(fmt.Printf("Completing after %d iterations", z))
			break
		} else {
			previousVal = l
		}
	}

	return l, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
