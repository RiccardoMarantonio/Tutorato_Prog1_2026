package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

type Pair struct {
	Parola    string
	Frequenza int
}

func LeggiFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func EstraiParole(testo string) []string {
	testo = strings.ToLower(testo)
	var parole []string

	for _, r := range testo {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			continue
		}
		testo = strings.ReplaceAll(testo, string(r), " ")
	}
	fields := strings.Fields(testo)
	for _, parola := range fields {
		parolaPulita := ""
		for _, r := range parola {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				parolaPulita += string(r)
			}
		}
		if parolaPulita != "" {
			parole = append(parole, parolaPulita)
		}
	}

	return parole
}

func ContaFrequenze(parole []string) map[string]int {
	frequenze := make(map[string]int)
	for _, p := range parole {
		frequenze[p]++
	}
	return frequenze
}

func OrdinaPerFrequenza(frequenze map[string]int) []Pair {
	var coppie []Pair
	for parola, freq := range frequenze {
		coppie = append(coppie, Pair{Parola: parola, Frequenza: freq})
	}

	sort.Slice(coppie, func(i, j int) bool {
		if coppie[i].Frequenza == coppie[j].Frequenza {
			return coppie[i].Parola < coppie[j].Parola
		}
		return coppie[i].Frequenza > coppie[j].Frequenza
	})

	return coppie
}

func StampaIstogramma(paroleOrdinate []Pair) {
	const maxBarre = 50

	if len(paroleOrdinate) == 0 {
		return
	}

	maxFreq := paroleOrdinate[0].Frequenza

	for _, p := range paroleOrdinate {
		if p.Frequenza < 2 {
			continue
		}

		numBarre := (p.Frequenza * maxBarre) / maxFreq
		if numBarre == 0 {
			numBarre = 1
		}

		barre := strings.Repeat("#", numBarre)

		fmt.Printf("%-20s: %-50s %d\n", p.Parola, barre, p.Frequenza)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s <file>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	path := os.Args[1]

	testo, err := LeggiFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errore lettura file: %v\n", err)
		os.Exit(1)
	}

	parole := EstraiParole(testo)

	frequenze := ContaFrequenze(parole)

	ordinato := OrdinaPerFrequenza(frequenze)

	StampaIstogramma(ordinato)
}
