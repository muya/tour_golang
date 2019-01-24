package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Nairobi": Vertex{
		-1.3108692, 36.6973553,
	},
}

func main() {
	fmt.Println(m)
}
