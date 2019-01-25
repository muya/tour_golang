package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader
// solution from: https://gist.github.com/rndD/5ba3c7c9883de396a4ea
func (m MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
