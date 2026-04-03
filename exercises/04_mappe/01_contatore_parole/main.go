package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// ContaParole legge una slice di parole e restituisce una mappa parola -> conteggio.
func ContaParole(parole []string) map[string]int {
	// TODO: implementare
	return nil
}

func main() {
	// TODO:
	// 1. Leggi parole da stdin una per riga (riga vuota = fine input)
	// 2. Chiama ContaParole per ottenere i conteggi
	// 3. Ordina le chiavi della mappa e stampa "parola: conteggio" per ognuna
	//
	// Suggerimento: usa bufio.NewScanner(os.Stdin) per leggere riga per riga
	// Per ordinare le chiavi: mettile in una slice e usa sort.Strings
	// Poi usa fmt.Printf per stampare
	_ = bufio.NewScanner(os.Stdin)
	_ = sort.Strings
	_ = fmt.Println
}
