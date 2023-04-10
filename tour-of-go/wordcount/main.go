package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	words := strings.Fields(s)
	for _, w := range words {
		count := 0
		for _, c := range words {
			if c == w {
				count++
				m[w] = count
			}
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
