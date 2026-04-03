package main

import (
	"testing"
)

func TestCalcola(t *testing.T) {
	tests := []struct {
		name string
		a    int
		op   string
		b    int
		want string
	}{
		{"somma", 5, "+", 3, "8"},
		{"sottrazione", 10, "-", 4, "6"},
		{"moltiplicazione", 7, "*", 3, "21"},
		{"divisione", 15, "/", 3, "5"},
		{"divisione zero", 10, "/", 0, "Divisione per zero"},
		{"operatore invalid", 5, "%", 3, "Operatore non valido"},
		{"negativi", -5, "+", -3, "-8"},
		{"divisione troncata", 7, "/", 2, "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Calcola(tt.a, tt.op, tt.b)
			if got != tt.want {
				t.Errorf("Calcola(%d, %q, %d) = %q, want %q", tt.a, tt.op, tt.b, got, tt.want)
			}
		})
	}
}
