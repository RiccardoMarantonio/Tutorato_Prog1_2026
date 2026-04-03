package main

import "fmt"

// Scambia scambia i valori di due variabili intere usando i puntatori.
func Scambia(a, b *int) {
	// TODO: implementare
}

// StampaInfo stampa valore, indirizzo e valore dereferenziato.
func StampaInfo(valore int, puntatore *int) {
	// TODO: implementare
}

func main() {
	x := 5
	y := 10
	fmt.Printf("Prima dello scambio: x = %d, y = %d\n", x, y)
	fmt.Printf("Indirizzo di x: %p\n", &x)
	fmt.Printf("Indirizzo di y: %p\n", &y)

	Scambia(&x, &y)

	fmt.Printf("Dopo lo scambio: x = %d, y = %d\n", x, y)
	StampaInfo(x, &x)
}
