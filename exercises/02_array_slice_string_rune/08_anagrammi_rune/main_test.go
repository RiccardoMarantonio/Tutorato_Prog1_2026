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

func TestAnagrammiRune(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:     "anagrammi base",
			input:    "listen silent",
			contains: []string{"anagrammi: true"},
		},
		{
			name:     "non anagrammi",
			input:    "hello world",
			contains: []string{"anagrammi: false"},
		},
		{
			name:     "case insensitive",
			input:    "Anna nana",
			contains: []string{"anagrammi: true"},
		},
		{
			name:     "stessa stringa",
			input:    "ciao ciao",
			contains: []string{"anagrammi: true"},
		},
		{
			name:     "lunghezze diverse",
			input:    "abc abcd",
			contains: []string{"anagrammi: false"},
		},
		{
			name:     "vuote",
			input:    " ",
			contains: []string{"anagrammi: true"},
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
