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

func TestFibonacciMemo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "10",
			contains: []string{
				"Fibonacci(10) = 55",
				"Chiamate senza memo: 177",
			},
		},
		{
			name:  "zero",
			input: "0",
			contains: []string{
				"Fibonacci(0) = 0",
				"Chiamate senza memo: 1",
			},
		},
		{
			name:  "uno",
			input: "1",
			contains: []string{
				"Fibonacci(1) = 1",
				"Chiamate senza memo: 1",
			},
		},
		{
			name:  "cinque",
			input: "5",
			contains: []string{
				"Fibonacci(5) = 5",
				"Chiamate senza memo: 15",
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
