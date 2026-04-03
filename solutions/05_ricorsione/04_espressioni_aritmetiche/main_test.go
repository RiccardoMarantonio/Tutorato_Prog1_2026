package main

import "testing"

func TestValutaEspressione(t *testing.T) {
	tests := []struct {
		expr string
		want float64
	}{
		{"3 + 5", 8},
		{"(3 + 5) * 2", 16},
		{"10 / (2 + 3)", 2},
		{"((1 + 2) * (3 + 4))", 21},
		{"2 * 3 + 4", 10},
		{"2 + 3 * 4", 14},
		{"10 - 3 - 2", 5},
		{"100", 100},
	}
	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			got, err := ValutaEspressione(tt.expr)
			if err != nil {
				t.Fatalf("ValutaEspressione(%q) error: %v", tt.expr, err)
			}
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 0.01 {
				t.Errorf("ValutaEspressione(%q) = %.2f, want %.2f", tt.expr, got, tt.want)
			}
		})
	}
}

func TestValutaEspressioneDivisioneZero(t *testing.T) {
	_, err := ValutaEspressione("10 / 0")
	if err == nil {
		t.Error("ValutaEspressione(10 / 0): errore atteso, got nil")
	}
}

func TestContaOperatori(t *testing.T) {
	ops := ContaOperatori("3 + 5 * 2 - 1")
	if ops["+"] != 1 || ops["*"] != 1 || ops["-"] != 1 || ops["/"] != 0 {
		t.Errorf("ContaOperatori: %+v", ops)
	}
}

func TestProfonditaParentesi(t *testing.T) {
	tests := []struct {
		expr string
		want int
	}{
		{"3 + 5", 0},
		{"(3 + 5)", 1},
		{"((1 + 2) * 3)", 2},
		{"(((1)))", 3},
	}
	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			if got := ProfonditaParentesi(tt.expr); got != tt.want {
				t.Errorf("ProfonditaParentesi(%q) = %d, want %d", tt.expr, got, tt.want)
			}
		})
	}
}

func TestEValida(t *testing.T) {
	tests := []struct {
		expr string
		want bool
	}{
		{"3 + 5", true},
		{"(3 + 5)", true},
		{"((3 + 5)", false},
		{"3 +", false},
		{"3 ++ 5", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			if got := EValida(tt.expr); got != tt.want {
				t.Errorf("EValida(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}
