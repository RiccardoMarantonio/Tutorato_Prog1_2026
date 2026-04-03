package main

import (
	"fmt"
)

// Libro rappresenta un libro nella biblioteca.
type Libro struct {
	ISBN        string
	Titolo      string
	Autore      string
	Anno        int
	Genere      string
	Disponibile bool
}

// Prestito rappresenta un prestito di un libro.
type Prestito struct {
	ISBN             string
	Utente           string
	DataPrestito     string
	DataRestituzione string
}

// Biblioteca gestisce libri e prestiti.
type Biblioteca struct {
	Libri    map[string]*Libro
	Prestiti []Prestito
}

// NuovaBiblioteca crea una nuova biblioteca vuota.
func NuovaBiblioteca() *Biblioteca {
	return &Biblioteca{
		Libri:    make(map[string]*Libro),
		Prestiti: []Prestito{},
	}
}

// AggiungiLibro aggiunge un libro. Errore se ISBN esiste già.
func AggiungiLibro(b *Biblioteca, l Libro) error {
	// TODO: implementare
	return nil
}

// CercaLibro cerca libri per titolo o autore (case-insensitive).
func CercaLibro(b Biblioteca, termine string) []Libro {
	// TODO: implementare
	return nil
}

// RegistraPrestito registra un prestito. Errore se libro non esiste o non disponibile.
func RegistraPrestito(b *Biblioteca, isbn, utente, data string) error {
	// TODO: implementare
	return nil
}

// RegistraRestituzione registra una restituzione. Errore se libro non in prestito.
func RegistraRestituzione(b *Biblioteca, isbn, data string) error {
	// TODO: implementare
	return nil
}

// LibriPerGenere restituisce mappa genere -> conteggio libri.
func LibriPerGenere(b Biblioteca) map[string]int {
	// TODO: implementare
	return nil
}

// PrestitiAperti restituisce i prestiti non ancora chiusi.
func PrestitiAperti(b Biblioteca) []Prestito {
	// TODO: implementare
	return nil
}

// Statistiche restituisce statistiche sulla biblioteca.
func Statistiche(b Biblioteca) map[string]int {
	// TODO: implementare
	return nil
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
				fmt.Printf("%s %q → %s (dal %s)\n", p.ISBN, titolo, p.Utente, p.DataPrestito)
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
