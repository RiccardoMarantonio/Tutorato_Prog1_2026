package main

import (
	"fmt"
	"strings"
)

// Contatto rappresenta una persona nella rubrica.
type Contatto struct {
	Nome     string
	Cognome  string
	Telefono string
}

// Rubrica contiene una lista di contatti.
type Rubrica struct {
	Contatti []Contatto
}

// toLower converte una stringa in minuscolo (solo ASCII).
func toLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= 'A' && b <= 'Z' {
			b += 32
		}
		result[i] = b
	}
	return string(result)
}

// Aggiungi aggiunge un contatto alla rubrica. Non aggiunge duplicati per telefono.
func Aggiungi(rubrica *Rubrica, c Contatto) {
	// TODO: implementare
}

// CercaPerNome restituisce tutti i contatti con il nome specificato (case-insensitive).
func CercaPerNome(rubrica Rubrica, nome string) []Contatto {
	// TODO: implementare
	return nil
}

// Rimuovi rimuove il contatto con il telefono specificato. Restituisce true se trovato.
func Rimuovi(rubrica *Rubrica, telefono string) bool {
	// TODO: implementare
	return false
}

// StampaRubrica stampa tutti i contatti.
func StampaRubrica(rubrica Rubrica) {
	for _, c := range rubrica.Contatti {
		fmt.Printf("%s, %s - %s\n", c.Cognome, c.Nome, c.Telefono)
	}
}

func main() {
	rubrica := Rubrica{}
	for {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "FINE" {
			break
		}
		switch cmd {
		case "AGGIUNGI":
			var c Contatto
			fmt.Scan(&c.Nome, &c.Cognome, &c.Telefono)
			Aggiungi(&rubrica, c)
			fmt.Printf("Contatto aggiunto: %s %s\n", c.Nome, c.Cognome)
		case "CERCA":
			var nome string
			fmt.Scan(&nome)
			results := CercaPerNome(rubrica, nome)
			fmt.Printf("\nRisultati ricerca %q:\n", nome)
			for _, c := range results {
				fmt.Printf("%s, %s - %s\n", c.Cognome, c.Nome, c.Telefono)
			}
			fmt.Println()
		case "RIMUOVI":
			var tel string
			fmt.Scan(&tel)
			if Rimuovi(&rubrica, tel) {
				fmt.Printf("Contatto rimosso: %s\n", tel)
			} else {
				fmt.Printf("Contatto non trovato: %s\n", tel)
			}
		case "STAMPA":
			fmt.Println("\n=== Rubrica ===")
			StampaRubrica(rubrica)
			fmt.Println()
		}
	}
	_ = strings.ToLower // hint: puoi usare strings.ToLower se vuoi
}
