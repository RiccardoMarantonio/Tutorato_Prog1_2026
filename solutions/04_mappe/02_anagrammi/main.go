package main

import (
	"fmt"
)

func sortBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[j] < b[i] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
}

func toLowerStr(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= 'A' && b <= 'Z' {
			b += 32
		}
		result[i] = b
	}
	return string(result)
}

func FirmaAnagramma(parola string) string {
	lower := toLowerStr(parola)
	b := []byte(lower)
	sortBytes(b)
	return string(b)
}

func RaggruppaAnagrammi(parole []string) [][]string {
	gruppi := make(map[string][]string)
	for _, p := range parole {
		firma := FirmaAnagramma(p)
		gruppi[firma] = append(gruppi[firma], p)
	}
	var result [][]string
	for _, g := range gruppi {
		result = append(result, g)
	}
	return result
}

func ContaAnagrammi(parole []string) map[string]int {
	conta := make(map[string]int)
	for _, p := range parole {
		firma := FirmaAnagramma(p)
		conta[firma]++
	}
	return conta
}

func main() {
	var n int
	fmt.Scan(&n)
	parole := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&parole[i])
	}

	gruppi := RaggruppaAnagrammi(parole)
	fmt.Println("Gruppi di anagrammi:")
	for _, g := range gruppi {
		fmt.Println(g)
	}

	conta := ContaAnagrammi(parole)
	fmt.Println("\nFirme piu comuni:")
	for firma, count := range conta {
		fmt.Printf("%s: %d\n", firma, count)
	}
}
