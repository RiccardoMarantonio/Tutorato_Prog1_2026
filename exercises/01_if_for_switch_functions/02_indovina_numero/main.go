package main

import "fmt"

// Confronta restituisce un feedback sul tentativo rispetto al segreto.
func Confronta(tentativo, segreto int) string {
	// TODO: implementare
	return ""
}

func main() {
	// TODO: implementa il gioco "indovina il numero"
	//
	// Il numero segreto e 42.
	// In un ciclo infinito:
	//   1. Leggi un tentativo da stdin
	//   2. Se e fuori dal range 1-100, stampa "Inserisci un numero tra 1 e 100" e continua
	//   3. Incrementa il contatore dei tentativi
	//   4. Chiama Confronta(tentativo, 42) e stampa il risultato
	//   5. Se il feedback e "Indovinato!", stampa il totale tentativi e esci dal ciclo
	//
	// Suggerimento: usa "for { ... }" per il ciclo infinito e "break" per uscire
	_ = fmt.Scan // placeholder
}
