package main

import (
	"testing"
)

func TestConfrontaParola(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"troppo corto", "go", "Troppo corto"},
		{"troppo lungo", "python", "Troppo lungo"},
		{"giusto", "golang", "Hai indovinato!"},
		{"lunghezza giusta sbagliata", "banana", "Lunghezza giusta, riprova!"},
		{"corto", "c", "Troppo corto"},
		{"lungo", "javascript", "Troppo lungo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConfrontaParola(tt.input, "golang")
			if got != tt.want {
				t.Errorf("ConfrontaParola(%q, %q) = %q, want %q", tt.input, "golang", got, tt.want)
			}
		})
	}
}
