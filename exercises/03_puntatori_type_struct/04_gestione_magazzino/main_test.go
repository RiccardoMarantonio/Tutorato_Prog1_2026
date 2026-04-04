package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runExercise(input string) (string, error) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func TestGestioneMagazzino(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "flusso completo",
			input: "AGGIUNGI P001 Laptop 899.99 10\n" +
				"AGGIUNGI P002 Mouse 29.99 50\n" +
				"SCARICO P001 3\n" +
				"CARICO P002 20\n" +
				"VALORE\n" +
				"SOTTO_SCORTA 15\n" +
				"FINE",
			contains: []string{
				"Prodotto P001 aggiunto",
				"Prodotto P002 aggiunto",
				"Scarico P001: 3 pezzi",
				"Carico P002: 20 pezzi",
				"Valore totale magazzino",
				"Prodotti sotto scorta",
			},
		},
		{
			name: "cerca prodotto",
			input: "AGGIUNGI P001 Laptop 899.99 10\n" +
				"CERCA P001\n" +
				"FINE",
			contains: []string{
				"Prodotto P001 aggiunto",
				"Trovato: P001",
				"Laptop",
			},
		},
		{
			name: "duplicato codice",
			input: "AGGIUNGI P001 Laptop 899.99 10\n" +
				"AGGIUNGI P001 Laptop2 500 5\n" +
				"FINE",
			contains: []string{
				"Prodotto P001 aggiunto",
				"Errore",
			},
		},
		{
			name: "scarico eccessivo",
			input: "AGGIUNGI P001 Laptop 899.99 2\n" +
				"SCARICO P001 100\n" +
				"FINE",
			contains: []string{
				"Errore",
			},
		},
		{
			name: "prodotto inesistente",
			input: "CERCA P999\n" +
				"FINE",
			contains: []string{
				"Errore",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			for _, c := range tt.contains {
				if !strings.Contains(got, c) {
					t.Errorf("output non contiene %q:\n%s", c, got)
				}
			}
		})
	}
}
