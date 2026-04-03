package main

import "testing"

func TestSommaRicorsiva(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 15},
		{10, 55},
		{100, 5050},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SommaRicorsiva(tt.n); got != tt.want {
				t.Errorf("SommaRicorsiva(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
