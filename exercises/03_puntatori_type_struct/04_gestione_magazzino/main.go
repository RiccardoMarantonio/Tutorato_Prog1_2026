package main

import (
	"fmt"
)

// Prodotto rappresenta un prodotto in magazzino.
type Prodotto struct {
	Codice    string
	Nome      string
	Categoria string
	Prezzo    float64
	Quantita  int
}

// Movimento rappresenta un carico o scarico.
type Movimento struct {
	CodiceProdotto string
	Tipo           string
	Quantita       int
}

// Magazzino gestisce prodotti e movimenti.
type Magazzino struct {
	Prodotti  []Prodotto
	Movimenti []Movimento
}

// AggiungiProdotto aggiunge un prodotto. Errore se il codice esiste già.
func AggiungiProdotto(m *Magazzino, p Prodotto) error {
	// TODO: implementare
	return nil
}

// RegistraMovimento registra un movimento e aggiorna la quantità.
func RegistraMovimento(m *Magazzino, mov Movimento) error {
	// TODO: implementare
	return nil
}

// ValoreMagazzino restituisce il valore totale del magazzino.
func ValoreMagazzino(m Magazzino) float64 {
	// TODO: implementare
	return 0
}

// ProdottiSottoScorta restituisce prodotti con quantità < soglia.
func ProdottiSottoScorta(m Magazzino, soglia int) []Prodotto {
	// TODO: implementare
	return nil
}

// RiepilogoCategoria restituisce mappa categoria -> totale pezzi.
func RiepilogoCategoria(m Magazzino) map[string]int {
	// TODO: implementare
	return nil
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
			fmt.Scan(&p.Codice, &p.Nome, &p.Categoria, &p.Prezzo, &p.Quantita)
			if err := AggiungiProdotto(m, p); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Prodotto %s aggiunto\n", p.Codice)
			}
		case "CARICO":
			var cod string
			var qta int
			fmt.Scan(&cod, &qta)
			mov := Movimento{cod, "CARICO", qta}
			if err := RegistraMovimento(m, mov); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Carico %s: %d pezzi\n", cod, qta)
			}
		case "SCARICO":
			var cod string
			var qta int
			fmt.Scan(&cod, &qta)
			mov := Movimento{cod, "SCARICO", qta}
			if err := RegistraMovimento(m, mov); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Scarico %s: %d pezzi\n", cod, qta)
			}
		case "VALORE":
			fmt.Printf("\nValore totale magazzino: %.2f EUR\n\n", ValoreMagazzino(*m))
		case "SOTTO_SCORTA":
			var soglia int
			fmt.Scan(&soglia)
			sotto := ProdottiSottoScorta(*m, soglia)
			fmt.Printf("\nProdotti sotto scorta (soglia: %d):\n", soglia)
			if len(sotto) == 0 {
				fmt.Println("Nessuno")
			} else {
				for _, p := range sotto {
					fmt.Printf("%s - %s - Quantita: %d\n", p.Codice, p.Nome, p.Quantita)
				}
			}
			fmt.Println()
		case "RIEPILOGO":
			fmt.Println("\nRiepilogo per categoria:")
			riep := RiepilogoCategoria(*m)
			for cat, qta := range riep {
				fmt.Printf("%s: %d\n", cat, qta)
			}
			fmt.Println()
		}
	}
}
