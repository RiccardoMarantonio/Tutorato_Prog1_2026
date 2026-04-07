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

func TestMedieArray(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     string
		contains []string
	}{
		{
			name:  "base 1-10",
			input: "1 2 3 4 5 6 7 8 9 10",
			want:  "",
			contains: []string{
				"Media: 5.5",
				"Sopra la media: 5",
				"Sotto la media: 5",
			},
		},
		{
			name:  "tutti uguali",
			input: "5 5 5 5 5 5 5 5 5 5",
			want:  "",
			contains: []string{
				"Media: 5",
				"Sopra la media: 0",
				"Sotto la media: 0",
			},
		},
		{
			name:  "negativi",
			input: "-5 -4 -3 -2 -1 0 1 2 3 4",
			want:  "",
			contains: []string{
				"Media: -0.5",
				"Sopra la media: 5",
				"Sotto la media: 5",
			},
		},
		{
			name:  "10-100",
			input: "10 20 30 40 50 60 70 80 90 100",
			want:  "",
			contains: []string{
				"Media: 55",
				"Sopra la media: 5",
				"Sotto la media: 5",
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
