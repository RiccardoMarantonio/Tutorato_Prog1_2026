package main

import "testing"

func TestFattoriale(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"0!", 0, 1},
		{"1!", 1, 1},
		{"2!", 2, 2},
		{"3!", 3, 6},
		{"4!", 4, 24},
		{"5!", 5, 120},
		{"12!", 12, 479001600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fattoriale(tt.n); got != tt.expected {
				t.Errorf("Fattoriale(%d) = %d; expected %d", tt.n, got, tt.expected)
			}
		})
	}
}
