package main

import "fmt"

func InvertiSlice(slice []int) []int {
	result := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		result[i] = slice[len(slice)-1-i]
	}
	return result
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

func EPalindromo(s string) bool {
	cleaned := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == ' ' {
			continue
		}
		cleaned = append(cleaned, toLower(b))
	}
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

func EPalindromoNumerico(numeri []int) bool {
	n := len(numeri)
	for i := 0; i < n/2; i++ {
		if numeri[i] != numeri[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%q è un palindromo: %v\n", s, EPalindromo(s))

	var n int
	fmt.Scan(&n)
	numeri := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Printf("%v è un palindromo numerico: %v\n", numeri, EPalindromoNumerico(numeri))
	fmt.Printf("Slice invertita: %v\n", InvertiSlice(numeri))
}
