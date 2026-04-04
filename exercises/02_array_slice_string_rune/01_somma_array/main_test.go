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

func TestSommaArray(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains string
	}{
		{"somma", "1 2 3 4 5\nsomma\n", "15"},
		{"prodotto", "1 2 3 4 5\nprodotto\n", "120"},
		{"prodotto con zero", "1 2 0 4 5\nprodotto\n", "0"},
		{"op invalida", "1 2 3 4 5\ndividi\n", "Operazione non valida"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			if !strings.Contains(got, tt.contains) {
				t.Errorf("input %q: output non contiene %q\n%s", tt.input, tt.contains, got)
			}
		})
	}
}
