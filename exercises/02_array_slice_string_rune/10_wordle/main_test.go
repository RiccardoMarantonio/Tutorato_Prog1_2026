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

func TestWordle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "vittoria al primo tentativo",
			input: "crane",
			contains: []string{
				"Tentativo 1/6: crane",
				"V V V V V",
				"Hai indovinato",
			},
		},
		{
			name:  "vittoria al secondo tentativo",
			input: "trace\ncrane",
			contains: []string{
				"Tentativo 1/6: trace",
				"Tentativo 2/6: crane",
				"Hai indovinato",
			},
		},
		{
			name:  "sconfitta dopo 6 tentativi",
			input: "abcde\nfghij\nklmno\npqrst\nuvwxy\nzabcd",
			contains: []string{
				"Tentativo 6/6",
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
