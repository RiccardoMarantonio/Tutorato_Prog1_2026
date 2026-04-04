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

func TestRotazioneStringhe(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "ciao\n2",
			contains: []string{
				"Originale: ciao",
				"Ruotata:   aoci",
			},
		},
		{
			name:  "k zero",
			input: "ciao\n0",
			contains: []string{
				"Originale: ciao",
				"Ruotata:   ciao",
			},
		},
		{
			name:  "k maggiore",
			input: "abc\n5",
			contains: []string{
				"Originale: abc",
				"Ruotata:   cab",
			},
		},
		{
			name:  "stringa vuota",
			input: "\n3",
			contains: []string{
				"Originale: ",
				"Ruotata:   ",
			},
		},
		{
			name:  "singolo carattere",
			input: "a\n10",
			contains: []string{
				"Originale: a",
				"Ruotata:   a",
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
