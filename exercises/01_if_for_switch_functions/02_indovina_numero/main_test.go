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

func TestIndovinaNumero(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			"vince in 1",
			"42\n",
			[]string{"Indovinato!", "1 tentativi"},
		},
		{
			"vince in 3",
			"50\n25\n42\n",
			[]string{"Troppo alto", "Troppo basso", "Indovinato!", "3 tentativi"},
		},
		{
			"fuori range",
			"0\n42\n",
			[]string{"Inserisci un numero tra 1 e 100", "Indovinato!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("output non contiene %q:\n%s", want, got)
				}
			}
		})
	}
}
