package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ValutaPN valuta un'espressione aritmetica in notazione prefissa
// (Notazione Polacca) implementando la logica di valutazione ricorsivamente.
func ValutaPN(tokens []string) (int, error) {
	// TODO: implementare la chiamata alla funzione ricorsiva iniziale
	return 0, errors.New("non implementato")
}

// valutaRicorsivo è la funzione helper ricorsiva suggerita.
// Restituisce il valore calcolato, i restanti token non ancora consumati, e un eventuale errore.
func valutaRicorsivo(tokens []string) (int, []string, error) {
	// TODO: implementare il caso base e le chiamate ricorsive (che "consumano" l'espressione da sinistra verso destra)
	return 0, nil, errors.New("non implementato")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		tokens := strings.Fields(input)

		if len(tokens) == 0 {
			return
		}

		risultato, err := ValutaPN(tokens)
		if err != nil {
			fmt.Printf("Errore: %v\n", err)
		} else {
			fmt.Printf("Risultato: %d\n", risultato)
		}
	}
}
