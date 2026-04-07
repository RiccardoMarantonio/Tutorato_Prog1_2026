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

func TestCifrarioCesare(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "ciao k3",
			input: "3\nfldr",
			contains: []string{
				"Decifrato: ciao",
			},
		},
		{
			name:  "abc k1",
			input: "1\nbcd",
			contains: []string{
				"Decifrato: abc",
			},
		},
		{
			name:  "abc uvw k3",
			input: "3\ndef ghil\nabc xyz",
			contains: []string{
				"Decifrato: abc def",
				"xyz uvw",
			},
		},
		{
			name:  "k zero",
			input: "0\nhello",
			contains: []string{
				"Decifrato: hello",
			},
		},
		{
			name:  "k grande",
			input: "28\ncde",
			contains: []string{
				"Decifrato: abc",
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
