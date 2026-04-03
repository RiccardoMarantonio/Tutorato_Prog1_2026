package main

import (
	"fmt"
)

type Prodotto struct {
	Codice    string
	Nome      string
	Categoria string
	Prezzo    float64
	Quantita  int
}

type Movimento struct {
	CodiceProdotto string
	Tipo           string
	Quantita       int
}

type Magazzino struct {
	Prodotti  []Prodotto
	Movimenti []Movimento
}

func AggiungiProdotto(m *Magazzino, p Prodotto) error {
	for _, prod := range m.Prodotti {
		if prod.Codice == p.Codice {
			return fmt.Errorf("prodotto %s gia esistente", p.Codice)
		}
	}
	m.Prodotti = append(m.Prodotti, p)
	return nil
}

func RegistraMovimento(m *Magazzino, mov Movimento) error {
	if mov.Quantita <= 0 {
		return fmt.Errorf("quantita deve essere maggiore di zero")
	}
	for i := range m.Prodotti {
		if m.Prodotti[i].Codice == mov.CodiceProdotto {
			if mov.Tipo == "CARICO" {
				m.Prodotti[i].Quantita += mov.Quantita
				m.Movimenti = append(m.Movimenti, mov)
				return nil
			} else if mov.Tipo == "SCARICO" {
				if m.Prodotti[i].Quantita < mov.Quantita {
					return fmt.Errorf("quantita insufficiente per %s (disponibili: %d, richiesti: %d)",
						mov.CodiceProdotto, m.Prodotti[i].Quantita, mov.Quantita)
				}
				m.Prodotti[i].Quantita -= mov.Quantita
				m.Movimenti = append(m.Movimenti, mov)
				return nil
			}
			return fmt.Errorf("tipo movimento non valido: %s", mov.Tipo)
		}
	}
	return fmt.Errorf("prodotto %s non trovato", mov.CodiceProdotto)
}

func ValoreMagazzino(m Magazzino) float64 {
	totale := 0.0
	for _, p := range m.Prodotti {
		totale += p.Prezzo * float64(p.Quantita)
	}
	return totale
}

func ProdottiSottoScorta(m Magazzino, soglia int) []Prodotto {
	var sotto []Prodotto
	for _, p := range m.Prodotti {
		if p.Quantita < soglia {
			sotto = append(sotto, p)
		}
	}
	return sotto
}

func RiepilogoCategoria(m Magazzino) map[string]int {
	result := make(map[string]int)
	for _, p := range m.Prodotti {
		result[p.Categoria] += p.Quantita
	}
	return result
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
