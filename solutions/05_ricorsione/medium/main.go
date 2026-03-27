package main

import (
	"fmt"
	"strings"
)

func isPalRunes(r []rune) bool {
	if len(r) <= 1 {
		return true
	}
	if r[0] != r[len(r)-1] {
		return false
	}
	return isPalRunes(r[1 : len(r)-1])
}

func IsPalindrome(s string) bool {
	normalized := strings.ReplaceAll(strings.ToLower(s), " ", "")
	return isPalRunes([]rune(normalized))
}

func main() {
	fmt.Println(IsPalindrome("i topi non avevano nipoti"))
}
