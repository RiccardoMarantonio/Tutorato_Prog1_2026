package main

import "fmt"

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
	for _, prod := range m.Prodotti {
		if prod.Codice == p.Codice {
			return fmt.Errorf("prodotto %s gia esistente", p.Codice)
		}
	}
	m.Prodotti = append(m.Prodotti, p)
	return nil
}

func RegistraCarico(m *Magazzino, codice string, qta int) error {
	if qta <= 0 {
		return fmt.Errorf("quantita deve essere maggiore di zero")
	}
	for i := range m.Prodotti {
		if m.Prodotti[i].Codice == codice {
			m.Prodotti[i].Quantita += qta
			return nil
		}
	}
	return fmt.Errorf("prodotto %s non trovato", codice)
}

func RegistraScarico(m *Magazzino, codice string, qta int) error {
	if qta <= 0 {
		return fmt.Errorf("quantita deve essere maggiore di zero")
	}
	for i := range m.Prodotti {
		if m.Prodotti[i].Codice == codice {
			if m.Prodotti[i].Quantita < qta {
				return fmt.Errorf("quantita insufficiente per %s (disponibili: %d, richiesti: %d)",
					codice, m.Prodotti[i].Quantita, qta)
			}
			m.Prodotti[i].Quantita -= qta
			return nil
		}
	}
	return fmt.Errorf("prodotto %s non trovato", codice)
}

func ValoreMagazzino(m Magazzino) float64 {
	totale := 0.0
	for _, p := range m.Prodotti {
		totale += p.Prezzo * float64(p.Quantita)
	}
	return totale
}

func ProdottiSottoScorta(m Magazzino, soglia int) int {
	count := 0
	for _, p := range m.Prodotti {
		if p.Quantita < soglia {
			count++
		}
	}
	return count
}

func CercaProdotto(m Magazzino, codice string) (Prodotto, error) {
	for _, p := range m.Prodotti {
		if p.Codice == codice {
			return p, nil
		}
	}
	return Prodotto{}, fmt.Errorf("prodotto %s non trovato", codice)
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
