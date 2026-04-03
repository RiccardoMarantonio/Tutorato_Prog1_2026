package main

import (
	"fmt"
)

// IsMultiplo restituisce true se n è multiplo di divisore.
func IsMultiplo(n, divisore int) bool {
	// TODO: implementare
	return false
}

// IsPrimo restituisce true se n è un numero primo.
func IsPrimo(n int) bool {
	// TODO: implementare
	return false
}

// FizzBuzz restituisce la stringa corretta per n secondo le regole FizzBuzz+Prime.
func FizzBuzz(n int) string {
	// TODO: implementare
	return ""
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
