package main

import (
	"fmt"
	"strings"
)

type Contatto struct {
	Nome     string
	Cognome  string
	Telefono string
}

type Rubrica struct {
	Contatti []Contatto
}

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

func Aggiungi(rubrica *Rubrica, c Contatto) {
	for _, existing := range rubrica.Contatti {
		if existing.Telefono == c.Telefono {
			return
		}
	}
	rubrica.Contatti = append(rubrica.Contatti, c)
}

func CercaPerNome(rubrica Rubrica, nome string) []Contatto {
	var results []Contatto
	nomeLower := toLower(nome)
	for _, c := range rubrica.Contatti {
		if toLower(c.Nome) == nomeLower {
			results = append(results, c)
		}
	}
	return results
}

func Rimuovi(rubrica *Rubrica, telefono string) bool {
	for i, c := range rubrica.Contatti {
		if c.Telefono == telefono {
			rubrica.Contatti = append(rubrica.Contatti[:i], rubrica.Contatti[i+1:]...)
			return true
		}
	}
	return false
}

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
	_ = strings.ToLower
}
