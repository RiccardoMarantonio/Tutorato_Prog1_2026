package main

import "testing"

func TestConfronta(t *testing.T) {
	tests := []struct {
		name      string
		tentativo int
		segreto   int
		want      string
	}{
		{"troppo alto", 50, 42, "Troppo alto"},
		{"troppo basso", 25, 42, "Troppo basso"},
		{"indovinato", 42, 42, "Indovinato!"},
		{"bordo alto", 100, 42, "Troppo alto"},
		{"bordo basso", 1, 42, "Troppo basso"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Confronta(tt.tentativo, tt.segreto); got != tt.want {
				t.Errorf("Confronta(%d, %d) = %q, want %q", tt.tentativo, tt.segreto, got, tt.want)
			}
		})
	}
}

func TestContaTentativi(t *testing.T) {
	tests := []struct {
		name    string
		storico []int
		want    int
	}{
		{"vuoto", []int{}, 0},
		{"uno", []int{42}, 1},
		{"quattro", []int{50, 25, 35, 42}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContaTentativi(tt.storico); got != tt.want {
				t.Errorf("ContaTentativi(%v) = %d, want %d", tt.storico, got, tt.want)
			}
		})
	}
}
