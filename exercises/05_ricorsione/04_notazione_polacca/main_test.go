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

func TestValutaPN(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Semplice addizione",
			input:       "+ 2 3",
			expected:    "Risultato: 5",
			expectError: false,
		},
		{
			name:        "Sottrazione negativa",
			input:       "- 10 15",
			expected:    "Risultato: -5",
			expectError: false,
		},
		{
			name:        "Espressione complessa",
			input:       "- + 5 * 1 2 3",
			expected:    "Risultato: 4",
			expectError: false,
		},
		{
			name:        "Divisione per zero",
			input:       "/ 10 0",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Mancanza di operandi",
			input:       "+ 4",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Token extra non consumati",
			input:       "+ 2 3 4",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Input non numerico",
			input:       "+ a 2",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := runExercise(tc.input)
			if err != nil {
				// Il programma go run . esce con stato di errore? Normalizza.
				// Noi non facciamo log.Fatal ma solo fmt.Printf("Errore: %v\n", err)
				// quindi l'uscita è di successo, ma l'output inizierà per "Errore:"
			}

			if tc.expectError {
				if !strings.HasPrefix(got, "Errore:") {
					t.Errorf("Atteso un messaggio di errore (che inizia con 'Errore:'), ottenuto: %q", got)
				}
			} else {
				if got != tc.expected {
					t.Errorf("Risultato errato, atteso %q, ottenuto %q", tc.expected, got)
				}
			}
		})
	}
}
