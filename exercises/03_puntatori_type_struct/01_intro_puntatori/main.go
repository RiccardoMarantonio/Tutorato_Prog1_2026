package main

import "fmt"

// Scambia scambia i valori di due variabili intere usando i puntatori.
func Scambia(a, b *int) {
	// TODO: implementare
}

// Punto rappresenta un punto nel piano cartesiano.
type Punto struct {
	X float64
	Y float64
}

// Sqrt approssimato con il metodo di Newton.
func Sqrt(x float64) float64 {
	// TODO: implementare
	return 0
}

// DistanzaDaOrigine calcola la distanza del punto dall'origine.
func DistanzaDaOrigine(p Punto) float64 {
	// TODO: implementare
	return 0
}

// Sposta sposta il punto di dx e dy. Modifica la struct originale.
func Sposta(p *Punto, dx, dy float64) {
	// TODO: implementare
}

func main() {
	// TODO:
	// Parte 1 - Puntatori:
	//   1. Crea due variabili x=5, y=10
	//   2. Stampa "=== Puntatori ==="
	//   3. Stampa "Prima: x = 5, y = 10"
	//   4. Chiama Scambia(&x, &y)
	//   5. Stampa "Dopo:  x = 10, y = 5"
	//
	// Parte 2 - Struct Punto:
	//   1. Stampa "=== Struct Punto ==="
	//   2. Leggi px, py da stdin
	//   3. Crea un Punto e stampa "Punto: (px, py)"
	//   4. Stampa "Distanza dall'origine: X"
	//   5. Leggi dx, dy da stdin
	//   6. Chiama Sposta(&punto, dx, dy)
	//   7. Stampa "Dopo spostamento (dx, dy): (nuovaX, nuovaY)"
	_ = fmt.Scan // placeholder
}
