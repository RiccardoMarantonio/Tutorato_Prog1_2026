package main

import (
	"bufio"
	"fmt"
	"os"
)

func toLowerRune(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func sortRunes(runes []rune) {
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
}

func FirmaRune(s string) []rune {
	r := []rune(s)
	for i := range r {
		r[i] = toLowerRune(r[i])
	}
	sortRunes(r)
	return r
}

func SonoAnagrammi(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	f1 := FirmaRune(s1)
	f2 := FirmaRune(s2)
	for i := range f1 {
		if f1[i] != f2[i] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var s1, s2 string
	fmt.Sscanf(line, "%s %s", &s1, &s2)
	fmt.Printf("%q e %q sono anagrammi: %v\n", s1, s2, SonoAnagrammi(s1, s2))
}
