package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toLowerByte(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

func EPalindromo(s string) bool {
	cleaned := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		cleaned = append(cleaned, toLowerByte(s[i]))
	}
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("%q e un palindromo: %v\n", s, EPalindromo(s))
	_ = strings.ToLower
}
