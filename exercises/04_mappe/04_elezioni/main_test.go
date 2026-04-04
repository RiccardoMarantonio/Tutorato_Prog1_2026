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

func TestElezioni(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "flusso completo",
			input: "CANDIDATI Rossi Bianchi Verdi\n" +
				"SEZIONE Centro 1000\n" +
				"SEZIONE Nord 800\n" +
				"VOTO Centro Rossi\n" +
				"VOTO Centro Bianchi\n" +
				"VOTO Nord Verdi\n" +
				"RISULTATI\n" +
				"AFFLUENZA\n" +
				"VINCITORE\n" +
				"FINE",
			contains: []string{
				"Candidati registrati",
				"Sezione Centro creata",
				"Sezione Nord creata",
				"Voto registrato: Centro",
				"Voto registrato: Nord",
				"Risultati Nazionali",
				"Rossi: 1",
				"Bianchi: 1",
				"Verdi: 1",
				"Affluenza",
				"Vincitore",
			},
		},
		{
			name: "ballottaggio",
			input: "CANDIDATI Rossi Bianchi Verdi\n" +
				"SEZIONE Centro 1000\n" +
				"VOTO Centro Rossi\n" +
				"VOTO Centro Rossi\n" +
				"VOTO Centro Bianchi\n" +
				"BALLOTTAGGIO\n" +
				"FINE",
			contains: []string{
				"Ballottaggio",
			},
		},
		{
			name: "voto sezione inesistente",
			input: "CANDIDATI Rossi\n" +
				"VOTO Centro Rossi\n" +
				"FINE",
			contains: []string{
				"Errore",
			},
		},
		{
			name: "voto candidato non valido",
			input: "CANDIDATI Rossi\n" +
				"SEZIONE Centro 100\n" +
				"VOTO Centro Bianchi\n" +
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
