package main

import "fmt"

type Rettangolo struct {
	Base    float64
	Altezza float64
}

func Area(r Rettangolo) float64 {
	return r.Base * r.Altezza
}

func Perimetro(r Rettangolo) float64 {
	return 2 * (r.Base + r.Altezza)
}

func Scala(r *Rettangolo, fattore float64) {
	r.Base *= fattore
	r.Altezza *= fattore
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
