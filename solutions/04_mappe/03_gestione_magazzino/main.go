package main

import "fmt"

type Prodotto struct {
	Nome     string
	Prezzo   float64
	Quantita int
}

type Magazzino map[string]Prodotto

func AggiungiProdotto(m Magazzino, codice string, p Prodotto) error {
	if _, ok := m[codice]; ok {
		return fmt.Errorf("prodotto %s gia esistente", codice)
	}
	m[codice] = p
	return nil
}

func RegistraCarico(m Magazzino, codice string, qta int) error {
	if qta <= 0 {
		return fmt.Errorf("quantita deve essere maggiore di zero")
	}
	prod, ok := m[codice]
	if !ok {
		return fmt.Errorf("prodotto %s non trovato", codice)
	}
	prod.Quantita += qta
	m[codice] = prod
	return nil
}

func RegistraScarico(m Magazzino, codice string, qta int) error {
	if qta <= 0 {
		return fmt.Errorf("quantita deve essere maggiore di zero")
	}
	prod, ok := m[codice]
	if !ok {
		return fmt.Errorf("prodotto %s non trovato", codice)
	}
	if prod.Quantita < qta {
		return fmt.Errorf("quantita insufficiente per %s (disponibili: %d, richiesti: %d)",
			codice, prod.Quantita, qta)
	}
	prod.Quantita -= qta
	m[codice] = prod
	return nil
}

func ValoreMagazzino(m Magazzino) float64 {
	totale := 0.0
	for _, p := range m {
		totale += p.Prezzo * float64(p.Quantita)
	}
	return totale
}

func ProdottiSottoScorta(m Magazzino, soglia int) int {
	count := 0
	for _, p := range m {
		if p.Quantita < soglia {
			count++
		}
	}
	return count
}

func CercaProdotto(m Magazzino, codice string) (Prodotto, error) {
	prod, ok := m[codice]
	if !ok {
		return Prodotto{}, fmt.Errorf("prodotto %s non trovato", codice)
	}
	return prod, nil
}

func NumeroProdotti(m Magazzino) int {
	return len(m)
}

func main() {
	m := make(Magazzino)
	for {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "FINE" {
			break
		}
		switch cmd {
		case "AGGIUNGI":
			var codice string
			var p Prodotto
			fmt.Scan(&codice, &p.Nome, &p.Prezzo, &p.Quantita)
			if err := AggiungiProdotto(m, codice, p); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Prodotto %s aggiunto\n", codice)
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
			fmt.Printf("\nValore totale magazzino: %.2f EUR\n\n", ValoreMagazzino(m))
		case "SOTTO_SCORTA":
			var soglia int
			fmt.Scan(&soglia)
			count := ProdottiSottoScorta(m, soglia)
			fmt.Printf("\nProdotti sotto scorta (soglia: %d): %d\n\n", soglia, count)
		case "NUMERO":
			fmt.Printf("\nNumero prodotti: %d\n\n", NumeroProdotti(m))
		case "CERCA":
			var cod string
			fmt.Scan(&cod)
			p, err := CercaProdotto(m, cod)
			if err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Trovato: %s - %s - %.2f EUR - Quantita: %d\n", cod, p.Nome, p.Prezzo, p.Quantita)
			}
		}
	}
}
