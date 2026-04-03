package main

import (
	"fmt"
)

type Libro struct {
	ISBN        string
	Titolo      string
	Autore      string
	Anno        int
	Genere      string
	Disponibile bool
}

type Prestito struct {
	ISBN             string
	Utente           string
	DataPrestito     string
	DataRestituzione string
}

type Biblioteca struct {
	Libri    map[string]*Libro
	Prestiti []Prestito
}

func NuovaBiblioteca() *Biblioteca {
	return &Biblioteca{
		Libri:    make(map[string]*Libro),
		Prestiti: []Prestito{},
	}
}

func toLowerByte(s string) string {
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

func AggiungiLibro(b *Biblioteca, l Libro) error {
	if _, ok := b.Libri[l.ISBN]; ok {
		return fmt.Errorf("libro con ISBN %s gia esistente", l.ISBN)
	}
	copia := l
	b.Libri[l.ISBN] = &copia
	return nil
}

func CercaLibro(b Biblioteca, termine string) []Libro {
	term := toLowerByte(termine)
	var results []Libro
	for _, l := range b.Libri {
		if toLowerByte(l.Titolo) == term || toLowerByte(l.Autore) == term ||
			toLowerByte(l.Titolo) == term || toLowerByte(l.Autore) == term {
			results = append(results, *l)
			continue
		}
		if contains(toLowerByte(l.Titolo), term) || contains(toLowerByte(l.Autore), term) {
			results = append(results, *l)
		}
	}
	return results
}

func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(substr) > len(s) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func RegistraPrestito(b *Biblioteca, isbn, utente, data string) error {
	libro, ok := b.Libri[isbn]
	if !ok {
		return fmt.Errorf("libro %s non trovato", isbn)
	}
	if !libro.Disponibile {
		return fmt.Errorf("libro %s non disponibile", isbn)
	}
	libro.Disponibile = false
	b.Prestiti = append(b.Prestiti, Prestito{
		ISBN:         isbn,
		Utente:       utente,
		DataPrestito: data,
	})
	return nil
}

func RegistraRestituzione(b *Biblioteca, isbn, data string) error {
	libro, ok := b.Libri[isbn]
	if !ok {
		return fmt.Errorf("libro %s non trovato", isbn)
	}
	if libro.Disponibile {
		return fmt.Errorf("libro %s non e in prestito", isbn)
	}
	libro.Disponibile = true
	for i := len(b.Prestiti) - 1; i >= 0; i-- {
		if b.Prestiti[i].ISBN == isbn && b.Prestiti[i].DataRestituzione == "" {
			b.Prestiti[i].DataRestituzione = data
			break
		}
	}
	return nil
}

func LibriPerGenere(b Biblioteca) map[string]int {
	gen := make(map[string]int)
	for _, l := range b.Libri {
		gen[l.Genere]++
	}
	return gen
}

func PrestitiAperti(b Biblioteca) []Prestito {
	var aperti []Prestito
	for _, p := range b.Prestiti {
		if p.DataRestituzione == "" {
			aperti = append(aperti, p)
		}
	}
	return aperti
}

func Statistiche(b Biblioteca) map[string]int {
	stats := map[string]int{
		"totale":      len(b.Libri),
		"disponibili": 0,
		"in_prestito": 0,
		"generi":      0,
	}
	gen := make(map[string]bool)
	for _, l := range b.Libri {
		if l.Disponibile {
			stats["disponibili"]++
		} else {
			stats["in_prestito"]++
		}
		gen[l.Genere] = true
	}
	stats["generi"] = len(gen)
	return stats
}

func main() {
	b := NuovaBiblioteca()
	for {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "FINE" {
			break
		}
		switch cmd {
		case "AGGIUNGI":
			var l Libro
			fmt.Scan(&l.ISBN, &l.Titolo, &l.Autore, &l.Anno, &l.Genere)
			l.Disponibile = true
			if err := AggiungiLibro(b, l); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Libro aggiunto: %s\n", l.Titolo)
			}
		case "CERCA":
			var termine string
			fmt.Scan(&termine)
			results := CercaLibro(*b, termine)
			fmt.Printf("\nRisultati ricerca %q:\n", termine)
			for _, l := range results {
				fmt.Printf("- %s (%s, %d) [%s]\n", l.Titolo, l.Autore, l.Anno, l.Genere)
			}
			fmt.Println()
		case "PRESTITO":
			var isbn, utente, data string
			fmt.Scan(&isbn, &utente, &data)
			if err := RegistraPrestito(b, isbn, utente, data); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Prestito registrato: %s a %s\n", isbn, utente)
			}
		case "RESTITUZIONE":
			var isbn, data string
			fmt.Scan(&isbn, &data)
			if err := RegistraRestituzione(b, isbn, data); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Restituzione registrata: %s\n", isbn)
			}
		case "PER_GENERI":
			fmt.Println("\nLibri per genere:")
			for gen, count := range LibriPerGenere(*b) {
				fmt.Printf("%s: %d\n", gen, count)
			}
			fmt.Println()
		case "PRESTITI_APERTI":
			fmt.Println("\nPrestiti aperti:")
			for _, p := range PrestitiAperti(*b) {
				libro := b.Libri[p.ISBN]
				titolo := "Sconosciuto"
				if libro != nil {
					titolo = libro.Titolo
				}
				fmt.Printf("%s %q -> %s (dal %s)\n", p.ISBN, titolo, p.Utente, p.DataPrestito)
			}
			fmt.Println()
		case "STATISTICHE":
			fmt.Println("\nStatistiche:")
			stats := Statistiche(*b)
			fmt.Printf("Totale libri: %d\n", stats["totale"])
			fmt.Printf("Disponibili: %d\n", stats["disponibili"])
			fmt.Printf("In prestito: %d\n", stats["in_prestito"])
			fmt.Printf("Generi diversi: %d\n", stats["generi"])
			fmt.Println()
		case "CATALOGO":
			fmt.Println("\nCatalogo completo:")
			for _, l := range b.Libri {
				stato := "DISPONIBILE"
				if !l.Disponibile {
					stato = "IN PRESTITO"
				}
				fmt.Printf("ISBN: %s | %s | %s | %d | %s | %s\n",
					l.ISBN, l.Titolo, l.Autore, l.Anno, l.Genere, stato)
			}
			fmt.Println()
		}
	}
}
