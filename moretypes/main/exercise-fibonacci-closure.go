package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
	num1 := 1
	num2 := 0
	next := 0

	performCalc := func() {
		next = num1 + num2
		num2 = num1
		num1 = next
	}

	return func() int {
		ret := num2

		performCalc()

		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
