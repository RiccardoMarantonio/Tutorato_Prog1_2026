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

func TestMaxMinSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "5\n10 3 7 1 9",
			contains: []string{
				"Max: 10",
				"Min: 1",
			},
		},
		{
			name:  "negativi",
			input: "4\n-5 -2 -9 -1",
			contains: []string{
				"Max: -1",
				"Min: -9",
			},
		},
		{
			name:  "singolo",
			input: "1\n42",
			contains: []string{
				"Max: 42",
				"Min: 42",
			},
		},
		{
			name:  "duplicati",
			input: "3\n5 5 5",
			contains: []string{
				"Max: 5",
				"Min: 5",
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
