package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// ContaParole legge una slice di parole e restituisce una mappa parola -> conteggio.
func ContaParole(parole []string) map[string]int {
	// TODO: implementare
	return nil
}

func main() {
	var parole []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parole = append(parole, line)
	}

	contatori := ContaParole(parole)

	// Ordina le chiavi per output deterministico
	var chiavi []string
	for k := range contatori {
		chiavi = append(chiavi, k)
	}
	sort.Strings(chiavi)

	for _, k := range chiavi {
		fmt.Printf("%s: %d\n", k, contatori[k])
	}
}
