package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hola Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v error = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}