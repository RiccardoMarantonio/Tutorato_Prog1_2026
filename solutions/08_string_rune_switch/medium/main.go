package main

import (
	"fmt"
	"unicode"
)

func TokenKind(s string) string {
	if s == "" {
		return "empty"
	}
	allDigits := true
	allLetters := true
	for _, r := range s {
		if !unicode.IsDigit(r) {
			allDigits = false
		}
		if !unicode.IsLetter(r) {
			allLetters = false
		}
	}

	switch {
	case allDigits:
		return "int"
	case allLetters:
		return "word"
	default:
		return "mixed"
	}
}

func main() {
	fmt.Println(TokenKind("123"))
}
