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

func TestTorreHanoi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     string
		contains []string
	}{
		{
			name:  "un disco",
			input: "1",
			want:  "",
			contains: []string{
				"Torre di Hanoi (1 dischi)",
				"1. Sposta disco da A a C",
				"Totale mosse: 1",
			},
		},
		{
			name:  "tre dischi",
			input: "3",
			want:  "",
			contains: []string{
				"Torre di Hanoi (3 dischi)",
				"Totale mosse: 7",
			},
		},
		{
			name:  "due dischi",
			input: "2",
			want:  "",
			contains: []string{
				"Torre di Hanoi (2 dischi)",
				"Totale mosse: 3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			if tt.want != "" && got != tt.want {
				t.Errorf("input %q:\ngot:\n%s\nwant:\n%s", tt.input, got, tt.want)
			}
			for _, c := range tt.contains {
				if !strings.Contains(got, c) {
					t.Errorf("output non contiene %q:\n%s", c, got)
				}
			}
		})
	}
}
