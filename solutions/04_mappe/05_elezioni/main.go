package main

import (
	"fmt"
)

type Sezione struct {
	Nome     string
	Voti     map[string]int
	Elettori int
	Votanti  int
}

type Elezione struct {
	Candidati []string
	Sezioni   map[string]*Sezione
}

func RegistraVoto(e *Elezione, sezione, candidato string) error {
	s, ok := e.Sezioni[sezione]
	if !ok {
		return fmt.Errorf("sezione %s non esiste", sezione)
	}
	valido := false
	for _, c := range e.Candidati {
		if c == candidato {
			valido = true
			break
		}
	}
	if !valido {
		return fmt.Errorf("candidato %s non valido", candidato)
	}
	if s.Votanti >= s.Elettori {
		return fmt.Errorf("sezione %s ha raggiunto il massimo degli elettori", sezione)
	}
	s.Voti[candidato]++
	s.Votanti++
	return nil
}

func RisultatiNazionali(e Elezione) map[string]int {
	ris := make(map[string]int)
	for _, c := range e.Candidati {
		ris[c] = 0
	}
	for _, s := range e.Sezioni {
		for cand, voti := range s.Voti {
			ris[cand] += voti
		}
	}
	return ris
}

func Affluenza(e Elezione) map[string]float64 {
	aff := make(map[string]float64)
	for nome, s := range e.Sezioni {
		if s.Elettori > 0 {
			aff[nome] = float64(s.Votanti) / float64(s.Elettori) * 100
		}
	}
	return aff
}

func Vincitore(e Elezione) (string, error) {
	ris := RisultatiNazionali(e)
	vinc := ""
	maxVoti := -1
	for c, v := range ris {
		if v > maxVoti {
			maxVoti = v
			vinc = c
		}
	}
	if maxVoti <= 0 {
		return "", fmt.Errorf("nessun voto registrato")
	}
	return vinc, nil
}

func Ballottaggio(e Elezione) (string, string, error) {
	ris := RisultatiNazionali(e)
	type cv struct {
		c    string
		voti int
	}
	var lista []cv
	for c, v := range ris {
		lista = append(lista, cv{c, v})
	}
	for i := 0; i < len(lista); i++ {
		for j := i + 1; j < len(lista); j++ {
			if lista[j].voti > lista[i].voti {
				lista[i], lista[j] = lista[j], lista[i]
			}
		}
	}
	if len(lista) < 2 {
		return "", "", fmt.Errorf("meno di 2 candidati")
	}
	return lista[0].c, lista[1].c, nil
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
			var c string
			for {
				if _, err := fmt.Scan(&c); err != nil {
					break
				}
				switch c {
				case "SEZIONE", "VOTO", "RISULTATI", "AFFLUENZA", "VINCITORE", "BALLOTTAGGIO", "FINE":
					goto afterCandidati
				}
				e.Candidati = append(e.Candidati, c)
			}
		afterCandidati:
			fmt.Printf("Candidati registrati: %v\n", e.Candidati)
			if c == "SEZIONE" {
				var nome string
				var elettori int
				fmt.Scan(&nome, &elettori)
				e.Sezioni[nome] = &Sezione{Nome: nome, Voti: make(map[string]int), Elettori: elettori}
				fmt.Printf("Sezione %s creata (%d elettori)\n", nome, elettori)
			} else if c != "FINE" {
				continue
			}
			if c == "FINE" {
				break
			}
		case "SEZIONE":
			var nome string
			var elettori int
			fmt.Scan(&nome, &elettori)
			e.Sezioni[nome] = &Sezione{Nome: nome, Voti: make(map[string]int), Elettori: elettori}
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
