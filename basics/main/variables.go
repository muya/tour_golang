package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
	i = 39
	java = true
	fmt.Println(i, c, python, java)
}
