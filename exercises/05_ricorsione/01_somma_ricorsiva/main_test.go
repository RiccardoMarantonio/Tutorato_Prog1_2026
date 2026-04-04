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

func TestSommaRicorsiva(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"zero", "0", "Somma da 1 a 0: 0"},
		{"uno", "1", "Somma da 1 a 1: 1"},
		{"cinque", "5", "Somma da 1 a 5: 15"},
		{"dieci", "10", "Somma da 1 a 10: 55"},
		{"cento", "100", "Somma da 1 a 100: 5050"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			if got != tt.want {
				t.Errorf("input %q:\ngot: %q\nwant: %q", tt.input, got, tt.want)
			}
		})
	}
}
