package main

import (
	"fmt"
)

// FirmaAnagramma restituisce la firma di una parola (caratteri ordinati).
func FirmaAnagramma(parola string) string {
	// TODO: implementare
	return ""
}

// RaggruppaAnagrammi raggruppa le parole per anagrammi.
func RaggruppaAnagrammi(parole []string) [][]string {
	// TODO: implementare
	return nil
}

// ContaAnagrammi restituisce mappa firma -> conteggio parole.
func ContaAnagrammi(parole []string) map[string]int {
	// TODO: implementare
	return nil
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
	fmt.Println("\nFirme più comuni:")
	for firma, count := range conta {
		fmt.Printf("%s: %d\n", firma, count)
	}
}
