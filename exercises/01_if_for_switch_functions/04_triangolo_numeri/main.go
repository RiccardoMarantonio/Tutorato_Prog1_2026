package main

import "fmt"

// CalcolaNumeroRiga restituisce l'ultimo numero presente nella riga specificata.
func CalcolaNumeroRiga(riga int) int {
	// TODO: implementare
	return 0
}

// SommaTriangolo restituisce la somma di tutti i numeri nel triangolo di n righe.
func SommaTriangolo(n int) int {
	// TODO: implementare
	return 0
}

// StampaTriangolo stampa il triangolo formattato.
func StampaTriangolo(n int) {
	// TODO: implementare
}

func main() {
	var n int
	fmt.Scan(&n)
	StampaTriangolo(n)
	fmt.Printf("Ultimo numero riga %d: %d\n", n, CalcolaNumeroRiga(n))
	fmt.Printf("Somma totale: %d\n", SommaTriangolo(n))
}
