package main

import (
	"fmt"
)

type Studente struct {
	Nome string
	Voti []int
}

func AggiungiVoto(registro map[string]*Studente, nome string, voto int) error {
	if voto < 1 || voto > 30 {
		return fmt.Errorf("voto %d non valido (deve essere tra 1 e 30)", voto)
	}
	if _, ok := registro[nome]; !ok {
		registro[nome] = &Studente{Nome: nome, Voti: []int{}}
	}
	registro[nome].Voti = append(registro[nome].Voti, voto)
	return nil
}

func MediaStudente(registro map[string]*Studente, nome string) (float64, error) {
	s, ok := registro[nome]
	if !ok {
		return 0, fmt.Errorf("studente %s non trovato", nome)
	}
	if len(s.Voti) == 0 {
		return 0, fmt.Errorf("studente %s non ha voti", nome)
	}
	somma := 0
	for _, v := range s.Voti {
		somma += v
	}
	return float64(somma) / float64(len(s.Voti)), nil
}

func MigliorPeggiore(registro map[string]*Studente) (string, string, error) {
	if len(registro) == 0 {
		return "", "", fmt.Errorf("registro vuoto")
	}
	migliore := ""
	peggiore := ""
	mediaMigliore := -1.0
	mediaPeggiore := 31.0
	for nome, s := range registro {
		if len(s.Voti) == 0 {
			continue
		}
		somma := 0
		for _, v := range s.Voti {
			somma += v
		}
		media := float64(somma) / float64(len(s.Voti))
		if media > mediaMigliore {
			mediaMigliore = media
			migliore = nome
		}
		if media < mediaPeggiore {
			mediaPeggiore = media
			peggiore = nome
		}
	}
	if migliore == "" || peggiore == "" {
		return "", "", fmt.Errorf("nessuno studente con voti")
	}
	return migliore, peggiore, nil
}

func DistribuzioneVoti(registro map[string]*Studente) map[string]int {
	dist := map[string]int{
		"Insufficiente": 0,
		"Sufficiente":   0,
		"Buono":         0,
		"Ottimo":        0,
	}
	for _, s := range registro {
		for _, v := range s.Voti {
			switch {
			case v >= 1 && v <= 17:
				dist["Insufficiente"]++
			case v >= 18 && v <= 23:
				dist["Sufficiente"]++
			case v >= 24 && v <= 27:
				dist["Buono"]++
			case v >= 28 && v <= 30:
				dist["Ottimo"]++
			}
		}
	}
	return dist
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
