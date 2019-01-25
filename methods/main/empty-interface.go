package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 412
	describe(i)

	i = "hola"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
