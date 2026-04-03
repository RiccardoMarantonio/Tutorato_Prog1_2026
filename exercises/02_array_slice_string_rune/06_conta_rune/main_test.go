package main

import "testing"

func TestContaRune(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		wantLettere int
		wantCifre   int
		wantSpazi   int
		wantAltri   int
	}{
		{
			"base",
			"Ciao 123!",
			4, 3, 1, 1,
		},
		{
			"solo lettere",
			"abc",
			3, 0, 0, 0,
		},
		{
			"vuoto",
			"",
			0, 0, 0, 0,
		},
		{
			"accenti",
			"cafe\u0301",
			5, 0, 0, 0, // cafe + accent = 5 rune
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContaRune(tt.s)
			if got["lettere"] != tt.wantLettere {
				t.Errorf("lettere = %d, want %d", got["lettere"], tt.wantLettere)
			}
			if got["cifre"] != tt.wantCifre {
				t.Errorf("cifre = %d, want %d", got["cifre"], tt.wantCifre)
			}
			if got["spazi"] != tt.wantSpazi {
				t.Errorf("spazi = %d, want %d", got["spazi"], tt.wantSpazi)
			}
			if got["altri"] != tt.wantAltri {
				t.Errorf("altri = %d, want %d", got["altri"], tt.wantAltri)
			}
		})
	}
}
