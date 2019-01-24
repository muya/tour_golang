package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["respuesta"] = 42
	fmt.Println("The Value:", m["respuesta"])

	m["respuesta"] = 48
	fmt.Println("The Value:", m["respuesta"])

	delete(m, "respuesta")
	fmt.Println("The Value:", m["respuesta"])

	v, ok := m["respuesta"]
	fmt.Println("The Value:", v, "Present?", ok)
}
