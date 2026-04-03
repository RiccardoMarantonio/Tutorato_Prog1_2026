package main

import "fmt"

// Rettangolo rappresenta un rettangolo con base e altezza.
type Rettangolo struct {
	Base    float64
	Altezza float64
}

// Area restituisce l'area del rettangolo.
func Area(r Rettangolo) float64 {
	// TODO: implementare
	return 0
}

// Perimetro restituisce il perimetro del rettangolo.
func Perimetro(r Rettangolo) float64 {
	// TODO: implementare
	return 0
}

// Scala modifica il rettangolo moltiplicando base e altezza per fattore.
func Scala(r *Rettangolo, fattore float64) {
	// TODO: implementare
}

func main() {
	var r Rettangolo
	fmt.Scan(&r.Base, &r.Altezza)
	fmt.Printf("Rettangolo: Base=%.2f, Altezza=%.2f\n", r.Base, r.Altezza)
	fmt.Printf("Area: %.2f\n", Area(r))
	fmt.Printf("Perimetro: %.2f\n", Perimetro(r))

	var fattore float64
	fmt.Scan(&fattore)
	Scala(&r, fattore)
	fmt.Printf("Dopo scala x%.0f: Base=%.2f, Altezza=%.2f\n", fattore, r.Base, r.Altezza)
	fmt.Printf("Nuova area: %.2f\n", Area(r))
}
