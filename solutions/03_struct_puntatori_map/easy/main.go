package main

import "fmt"

func WordFreq(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	return m
}

func main() {
	fmt.Println(WordFreq([]string{"go", "go", "lang"}))
}
