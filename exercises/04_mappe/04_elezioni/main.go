package main

import (
	"fmt"
)

// Sezione rappresenta una sezione elettorale.
type Sezione struct {
	Nome     string
	Voti     map[string]int
	Elettori int
	Votanti  int
}

// Elezione rappresenta un'elezione con candidati e sezioni.
type Elezione struct {
	Candidati []string
	Sezioni   map[string]*Sezione
}

// RegistraVoto registra un voto. Errore se sezione/candidato non validi o sezione piena.
func RegistraVoto(e *Elezione, sezione, candidato string) error {
	// TODO: implementare
	return nil
}

// RisultatiNazionali restituisce il totale voti per candidato.
func RisultatiNazionali(e Elezione) map[string]int {
	// TODO: implementare
	return nil
}

// Affluenza restituisce la percentuale di affluenza per sezione.
func Affluenza(e Elezione) map[string]float64 {
	// TODO: implementare
	return nil
}

// Vincitore restituisce il candidato con più voti.
func Vincitore(e Elezione) (string, error) {
	// TODO: implementare
	return "", nil
}

// Ballottaggio restituisce i due candidati più votati.
func Ballottaggio(e Elezione) (string, string, error) {
	// TODO: implementare
	return "", "", nil
}

func main() {
	e := &Elezione{Sezioni: make(map[string]*Sezione)}
	for {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "FINE" {
			break
		}
		switch cmd {
		case "CANDIDATI":
			for {
				var c string
				fmt.Scan(&c)
				if c == "SEZIONE" || c == "VOTO" || c == "RISULTATI" || c == "AFFLUENZA" || c == "VINCITORE" || c == "BALLOTTAGGIO" || c == "FINE" {
					// Rimetti indietro leggendo il prossimo ciclo
					fmt.Printf("Candidati registrati: %v\n", e.Candidati)
					// Il prossimo cmd sarà processato dal ciclo esterno
					// Per semplicità, salviamo il prossimo comando
					// In realtà qui semplifichiamo: i candidati sono sulla stessa riga
					break
				}
				e.Candidati = append(e.Candidati, c)
			}
		case "SEZIONE":
			var nome string
			var elettori int
			fmt.Scan(&nome, &elettori)
			e.Sezioni[nome] = &Sezione{
				Nome:     nome,
				Voti:     make(map[string]int),
				Elettori: elettori,
			}
			fmt.Printf("Sezione %s creata (%d elettori)\n", nome, elettori)
		case "VOTO":
			var sez, cand string
			fmt.Scan(&sez, &cand)
			if err := RegistraVoto(e, sez, cand); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Voto registrato: %s → %s\n", sez, cand)
			}
		case "RISULTATI":
			fmt.Println("\n=== Risultati Nazionali ===")
			ris := RisultatiNazionali(*e)
			for _, c := range e.Candidati {
				fmt.Printf("%s: %d\n", c, ris[c])
			}
			fmt.Println()
		case "AFFLUENZA":
			fmt.Println("\n=== Affluenza ===")
			aff := Affluenza(*e)
			for nome, pct := range aff {
				fmt.Printf("%s: %.2f%%\n", nome, pct)
			}
			fmt.Println()
		case "VINCITORE":
			vinc, err := Vincitore(*e)
			if err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				ris := RisultatiNazionali(*e)
				fmt.Printf("Vincitore: %s (%d voti)\n", vinc, ris[vinc])
			}
		case "BALLOTTAGGIO":
			c1, c2, err := Ballottaggio(*e)
			if err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Ballottaggio: %s vs %s\n", c1, c2)
			}
		}
	}
}
