package main

import (
	"fmt"
)

// Studente rappresenta uno studente con i suoi voti.
type Studente struct {
	Nome string
	Voti []int
}

// AggiungiVoto aggiunge un voto a uno studente, creandolo se non esiste.
func AggiungiVoto(registro map[string]*Studente, nome string, voto int) error {
	// TODO: implementare
	return nil
}

// MediaStudente calcola la media dei voti di uno studente.
func MediaStudente(registro map[string]*Studente, nome string) (float64, error) {
	// TODO: implementare
	return 0, nil
}

// MigliorPeggiore restituisce i nomi dello studente migliore e peggiore.
func MigliorPeggiore(registro map[string]*Studente) (string, string, error) {
	// TODO: implementare
	return "", "", nil
}

// DistribuzioneVoti restituisce la distribuzione dei voti per fascia.
func DistribuzioneVoti(registro map[string]*Studente) map[string]int {
	// TODO: implementare
	return nil
}

func main() {
	registro := make(map[string]*Studente)
	for {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "FINE" {
			break
		}
		switch cmd {
		case "AGGIUNGI":
			var nome string
			var voto int
			fmt.Scan(&nome, &voto)
			if err := AggiungiVoto(registro, nome, voto); err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Voto aggiunto per %s: %d\n", nome, voto)
			}
		case "MEDIA":
			var nome string
			fmt.Scan(&nome)
			media, err := MediaStudente(registro, nome)
			if err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				fmt.Printf("Media di %s: %.2f\n", nome, media)
			}
		case "MIGLIOR_PEGGIORE":
			migliore, peggiore, err := MigliorPeggiore(registro)
			if err != nil {
				fmt.Printf("Errore: %v\n", err)
			} else {
				m1, _ := MediaStudente(registro, migliore)
				m2, _ := MediaStudente(registro, peggiore)
				fmt.Printf("Migliore: %s (%.2f)\nPeggiore: %s (%.2f)\n", migliore, m1, peggiore, m2)
			}
		case "DISTRIBUZIONE":
			dist := DistribuzioneVoti(registro)
			fmt.Println("\nDistribuzione voti:")
			for _, fascia := range []string{"Insufficiente", "Sufficiente", "Buono", "Ottimo"} {
				fmt.Printf("%s: %d\n", fascia, dist[fascia])
			}
			fmt.Println()
		}
	}
}
