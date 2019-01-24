package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Indian Ocean"] = Vertex{
		-1.4345, 98.2394,
	}
	fmt.Println(m["Indian Ocean"])
}
