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
