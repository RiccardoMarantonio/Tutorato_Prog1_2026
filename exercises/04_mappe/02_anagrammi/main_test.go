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

func TestAnagrammi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "base",
			input: "6\n" +
				"listen\n" +
				"silent\n" +
				"enlist\n" +
				"hello\n" +
				"world\n" +
				"drowl",
			contains: []string{
				"Gruppi di anagrammi",
				"listen",
				"silent",
				"enlist",
				"hello",
				"world",
				"drowl",
				"Firme",
			},
		},
		{
			name: "tutti anagrammi",
			input: "3\n" +
				"listen\n" +
				"silent\n" +
				"enlist",
			contains: []string{
				"Gruppi di anagrammi",
				"listen",
				"silent",
				"enlist",
			},
		},
		{
			name: "nessun anagramma",
			input: "3\n" +
				"abc\n" +
				"def\n" +
				"ghi",
			contains: []string{
				"Gruppi di anagrammi",
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
