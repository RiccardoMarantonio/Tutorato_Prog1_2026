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

func TestRubricaTelefonica(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "aggiungi e cerca",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"AGGIUNGI Luigi Bianchi 3339876543\n" +
				"AGGIUNGI Mario Verdi 3331112222\n" +
				"CERCA Mario\n" +
				"STAMPA\n" +
				"FINE",
			contains: []string{
				"Contatto aggiunto: Mario Rossi",
				"Contatto aggiunto: Luigi Bianchi",
				"Contatto aggiunto: Mario Verdi",
				"Risultati ricerca",
				"Rossi, Mario - 3331234567",
				"Verdi, Mario - 3331112222",
				"=== Rubrica ===",
			},
		},
		{
			name: "rimuovi contatto",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"AGGIUNGI Luigi Bianchi 3339876543\n" +
				"RIMUOVI 3339876543\n" +
				"STAMPA\n" +
				"FINE",
			contains: []string{
				"Contatto rimosso: 3339876543",
				"Rossi, Mario - 3331234567",
			},
		},
		{
			name: "duplicato telefono",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"AGGIUNGI Mario Verdi 3331234567\n" +
				"STAMPA\n" +
				"FINE",
			contains: []string{
				"Contatto aggiunto: Mario Rossi",
			},
		},
		{
			name: "ricerca case insensitive",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"CERCA mario\n" +
				"FINE",
			contains: []string{
				"Rossi, Mario - 3331234567",
			},
		},
		{
			name: "rimuovi inesistente",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"RIMUOVI 0000000000\n" +
				"FINE",
			contains: []string{
				"Contatto non trovato: 0000000000",
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
