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

func TestRettangoloStruct(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "4 6\n2",
			contains: []string{
				"Base=4.00",
				"Altezza=6.00",
				"Area: 24.00",
				"Perimetro: 20.00",
				"Dopo scala x2: Base=8.00",
				"Altezza=12.00",
				"Nuova area: 96.00",
			},
		},
		{
			name:  "quadrato",
			input: "5 5\n3",
			contains: []string{
				"Base=5.00",
				"Altezza=5.00",
				"Area: 25.00",
				"Perimetro: 20.00",
				"Dopo scala x3: Base=15.00",
				"Altezza=15.00",
				"Nuova area: 225.00",
			},
		},
		{
			name:  "fattore uno",
			input: "10 3\n1",
			contains: []string{
				"Area: 30.00",
				"Perimetro: 26.00",
				"Dopo scala x1: Base=10.00",
				"Altezza=3.00",
				"Nuova area: 30.00",
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
