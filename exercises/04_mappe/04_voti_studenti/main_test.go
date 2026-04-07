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

func TestVotiStudenti(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "flusso completo",
			input: "AGGIUNGI Mario 25\n" +
				"AGGIUNGI Mario 28\n" +
				"AGGIUNGI Luigi 18\n" +
				"AGGIUNGI Anna 30\n" +
				"AGGIUNGI Anna 30\n" +
				"MEDIA Mario\n" +
				"MIGLIOR_PEGGIORE\n" +
				"DISTRIBUZIONE\n" +
				"FINE",
			contains: []string{
				"Voto aggiunto per Mario: 25",
				"Voto aggiunto per Mario: 28",
				"Media di Mario: 26.50",
				"Migliore: Anna",
				"Peggiore: Luigi",
				"Distribuzione voti",
			},
		},
		{
			name: "voto non valido",
			input: "AGGIUNGI Mario 31\n" +
				"FINE",
			contains: []string{
				"Errore",
			},
		},
		{
			name: "media studente inesistente",
			input: "MEDIA Inesistente\n" +
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
