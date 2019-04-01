package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount is counter world
func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range arr {
		m[v] = m[v] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
