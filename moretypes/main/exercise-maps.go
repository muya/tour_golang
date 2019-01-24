package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	fields := strings.Fields(s)

	for i := 0; i < len(fields); i++ {
		currentWord := fields[i]
		_, exists := m[currentWord]
		if exists {
			m[currentWord] = m[currentWord] + 1
		} else {
			m[currentWord] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
