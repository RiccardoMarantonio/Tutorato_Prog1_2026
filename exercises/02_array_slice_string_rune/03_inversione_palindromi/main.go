package main

import "fmt"

// InvertiSlice restituisce una nuova slice invertita senza modificare l'originale.
func InvertiSlice(slice []int) []int {
	// TODO: implementare
	return nil
}

// toLower converte un byte ASCII in minuscolo.
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

// EPalindromo verifica se una stringa è un palindromo (case-insensitive, ignora spazi).
func EPalindromo(s string) bool {
	// TODO: implementare
	return false
}

// EPalindromoNumerico verifica se una slice è un palindromo numerico.
func EPalindromoNumerico(numeri []int) bool {
	// TODO: implementare
	return false
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
