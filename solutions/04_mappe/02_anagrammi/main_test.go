package main

import "testing"

func TestFirmaAnagramma(t *testing.T) {
	tests := []struct {
		parola string
		want   string
	}{
		{"listen", "eilnst"},
		{"silent", "eilnst"},
		{"hello", "ehllo"},
		{"world", "dlorw"},
		{"", ""},
		{"a", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.parola, func(t *testing.T) {
			if got := FirmaAnagramma(tt.parola); got != tt.want {
				t.Errorf("FirmaAnagramma(%q) = %q, want %q", tt.parola, got, tt.want)
			}
		})
	}
}

func TestRaggruppaAnagrammi(t *testing.T) {
	parole := []string{"listen", "silent", "enlist", "hello", "world"}
	gruppi := RaggruppaAnagrammi(parole)

	if len(gruppi) != 3 {
		t.Errorf("RaggruppaAnagrammi: %d gruppi, want 3", len(gruppi))
	}

	trovato := false
	for _, g := range gruppi {
		if len(g) == 3 {
			trovato = true
			break
		}
	}
	if !trovato {
		t.Error("Nessun gruppo con 3 anagrammi trovato")
	}
}

func TestContaAnagrammi(t *testing.T) {
	parole := []string{"listen", "silent", "hello"}
	conta := ContaAnagrammi(parole)

	if conta["eilnst"] != 2 {
		t.Errorf("ContaAnagrammi[eilnst] = %d, want 2", conta["eilnst"])
	}
	if conta["ehllo"] != 1 {
		t.Errorf("ContaAnagrammi[ehllo] = %d, want 1", conta["ehllo"])
	}
}
