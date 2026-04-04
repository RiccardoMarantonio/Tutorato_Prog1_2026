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

func TestFusioneSliceOrdinate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "5\n1 3 5 7 9\n6\n2 3 5 8 9 10",
			contains: []string{
				"Merge: [1 2 3 3 5 5 7 8 9 9 10]",
				"Senza duplicati: [1 2 3 5 7 8 9 10]",
				"Intersezione: [3 5 9]",
			},
		},
		{
			name:  "slice vuota",
			input: "0\n3\n1 2 3",
			contains: []string{
				"Merge: [1 2 3]",
				"Intersezione: []",
			},
		},
		{
			name:  "identiche",
			input: "3\n1 2 3\n3\n1 2 3",
			contains: []string{
				"Merge: [1 1 2 2 3 3]",
				"Senza duplicati: [1 2 3]",
				"Intersezione: [1 2 3]",
			},
		},
		{
			name:  "nessuna intersezione",
			input: "3\n1 2 3\n3\n4 5 6",
			contains: []string{
				"Merge: [1 2 3 4 5 6]",
				"Intersezione: []",
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
