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

func TestEspressioniAritmetiche(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "somma semplice",
			input: "3 + 5",
			contains: []string{
				"3 + 5 = 8.00",
				"+: 1",
				"Massima profondita parentesi: 0",
				"Espressione valida: true",
			},
		},
		{
			name:  "con parentesi",
			input: "(3 + 5) * 2",
			contains: []string{
				"(3 + 5) * 2 = 16.00",
				"+: 1",
				"*: 1",
				"Massima profondita parentesi: 1",
				"Espressione valida: true",
			},
		},
		{
			name:  "precedenza operatori",
			input: "2 + 3 * 4",
			contains: []string{
				"2 + 3 * 4 = 14.00",
			},
		},
		{
			name:  "parentesi annidate",
			input: "((1 + 2) * (3 + 4))",
			contains: []string{
				"= 21.00",
				"Massima profondita parentesi: 2",
			},
		},
		{
			name:  "espressione non valida",
			input: "((3 + 5)",
			contains: []string{
				"Espressione valida: false",
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
