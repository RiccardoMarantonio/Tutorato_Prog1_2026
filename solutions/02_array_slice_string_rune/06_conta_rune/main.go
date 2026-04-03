package main

import (
	"bufio"
	"fmt"
	"os"
)

func ContaRune(s string) map[string]int {
	count := map[string]int{
		"lettere": 0,
		"cifre":   0,
		"spazi":   0,
		"altri":   0,
	}
	for _, r := range s {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'):
			count["lettere"]++
		case r >= '0' && r <= '9':
			count["cifre"]++
		case r == ' ':
			count["spazi"]++
		default:
			count["altri"]++
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	count := ContaRune(s)
	fmt.Printf("Totale rune: %d\n", len([]rune(s)))
	fmt.Printf("Lettere: %d\n", count["lettere"])
	fmt.Printf("Cifre: %d\n", count["cifre"])
	fmt.Printf("Spazi: %d\n", count["spazi"])
	fmt.Printf("Altri: %d\n", count["altri"])
}
