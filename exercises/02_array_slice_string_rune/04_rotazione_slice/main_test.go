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

func TestRotazioneSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "5\n1 2 3 4 5\n2",
			contains: []string{
				"Originale: [1 2 3 4 5]",
				"Ruotata:   [4 5 1 2 3]",
			},
		},
		{
			name:  "k zero",
			input: "3\n1 2 3\n0",
			contains: []string{
				"Originale: [1 2 3]",
				"Ruotata:   [1 2 3]",
			},
		},
		{
			name:  "k maggiore di n",
			input: "3\n1 2 3\n5",
			contains: []string{
				"Originale: [1 2 3]",
				"Ruotata:   [2 3 1]",
			},
		},
		{
			name:  "singolo elemento",
			input: "1\n7\n10",
			contains: []string{
				"Originale: [7]",
				"Ruotata:   [7]",
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
