package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isAlpha(t byte) bool {
	return (t >= 65 && t <= 90) || (t >= 97 && t <= 122)
}

func rot13(t byte) byte {
	var proposedNewVal byte
	if t+13 <= 122 {
		// check for addition
		proposedNewVal = t + 13
	} else {
		// check for subtraction
		proposedNewVal = t - 13
	}

	if isAlpha(proposedNewVal) {
		return proposedNewVal
	}

	return t
}

func (m *rot13Reader) Read(t []byte) (int, error) {
	n, err := m.r.Read(t)

	for i := 0; i < n; i++ {
		orig := t[i]
		if isAlpha(t[i]) {
			t[i] = rot13(t[i])
		} else {
			t[i] = t[i]
		}
		newVal := t[i]
		fmt.Println(orig, newVal)
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
