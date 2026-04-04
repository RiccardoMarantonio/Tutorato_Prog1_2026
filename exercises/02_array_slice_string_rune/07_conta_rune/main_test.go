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

func TestContaRune(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "Ciao 123!",
			contains: []string{
				"Lettere: 4",
				"Cifre: 3",
				"Spazi: 1",
				"Altri: 1",
			},
		},
		{
			name:  "solo lettere",
			input: "abc",
			contains: []string{
				"Lettere: 3",
				"Cifre: 0",
				"Spazi: 0",
				"Altri: 0",
			},
		},
		{
			name:  "stringa vuota",
			input: "",
			contains: []string{
				"Lettere: 0",
				"Cifre: 0",
				"Spazi: 0",
				"Altri: 0",
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
