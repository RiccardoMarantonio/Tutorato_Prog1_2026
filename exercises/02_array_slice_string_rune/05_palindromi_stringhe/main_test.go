package main

import "testing"

func TestEPalindromo(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"anna", "Anna", true},
		{"otto", "otto", true},
		{"ciao", "ciao", false},
		{"vuoto", "", true},
		{"singolo", "a", true},
		{"con spazi", "i topi non avevano nipoti", true},
		{"non palindromo con spazi", "hello world", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EPalindromo(tt.s); got != tt.want {
				t.Errorf("EPalindromo(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
