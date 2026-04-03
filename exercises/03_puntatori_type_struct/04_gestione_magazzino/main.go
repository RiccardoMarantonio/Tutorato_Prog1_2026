package main

import (
	"fmt"
)

type Prodotto struct {
	Codice   string
	Nome     string
	Prezzo   float64
	Quantita int
}

type Magazzino struct {
	Prodotti []Prodotto
}

func AggiungiProdotto(m *Magazzino, p Prodotto) error {
	// TODO: implementare
	return nil
}

func RegistraCarico(m *Magazzino, codice string, qta int) error {
	// TODO: implementare
	return nil
}

func RegistraScarico(m *Magazzino, codice string, qta int) error {
	// TODO: implementare
	return nil
}

func ValoreMagazzino(m Magazzino) float64 {
	// TODO: implementare
	return 0
}

func ProdottiSottoScorta(m Magazzino, soglia int) int {
	// TODO: implementare
	return 0
}

func CercaProdotto(m Magazzino, codice string) (Prodotto, error) {
	// TODO: implementare
	return Prodotto{}, nil
}

func main() {
	m := &Magazzino{}
	for {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "FINE" {
			break
		}
		switch cmd {
		case "AGGIUNGI":
			var p Prodotto
			fmt.Scan(&p.Codice, &p.Nome, &p.Prezzo, &p.Quantita)
			if err := AggiungiProdotto(m, p); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Prodotto %s aggiunto\n", p.Codice)
			}
		case "CARICO":
			var cod string
			var qta int
			fmt.Scan(&cod, &qta)
			if err := RegistraCarico(m, cod, qta); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Carico %s: %d pezzi\n", cod, qta)
			}
		case "SCARICO":
			var cod string
			var qta int
			fmt.Scan(&cod, &qta)
			if err := RegistraScarico(m, cod, qta); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Scarico %s: %d pezzi\n", cod, qta)
			}
		case "VALORE":
			fmt.Printf("\nValore totale magazzino: %.2f EUR\n\n", ValoreMagazzino(*m))
		case "SOTTO_SCORTA":
			var soglia int
			fmt.Scan(&soglia)
			count := ProdottiSottoScorta(*m, soglia)
			fmt.Printf("\nProdotti sotto scorta (soglia: %d): %d\n\n", soglia, count)
		case "CERCA":
			var cod string
			fmt.Scan(&cod)
			p, err := CercaProdotto(*m, cod)
			if err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Trovato: %s - %s - %.2f EUR - Quantita: %d\n", p.Codice, p.Nome, p.Prezzo, p.Quantita)
			}
		}
	}
}
