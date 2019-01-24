package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//var t = [][]uint8{
	//	{uint8(1), uint8(1)},
	//}

	var b = make([][]uint8, dy)

	for c := 0; c < dy; c++ {
		b[c] = make([]uint8, dx)
	}

	for d := 0; d < dy; d++ {
		for e := 0; e < dx; e++ {
			//b[d][e] = uint8(((d - 1) / (e + 1)) * 2)
			//b[d][e] = uint8(d - e)
			//b[d][e] = uint8((d * d) - (e * e))
			//b[d][e] = uint8((d * d * d) - (e * e * e))
			b[d][e] = uint8((d * d * d) * (e * e * e))
			//b[d][e] = uint8((d * e) - (d * e) + 1)
		}
	}

	return b
}

func main() {
	pic.Show(Pic)
}
