package main

import "testing"

func TestSonoAnagrammi(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{"base", "listen", "silent", true},
		{"no", "hello", "world", false},
		{"case insensitive", "Anna", "nana", true},
		{"stessa stringa", "ciao", "ciao", true},
		{"lunghezze diverse", "abc", "abcd", false},
		{"vuote", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SonoAnagrammi(tt.s1, tt.s2); got != tt.want {
				t.Errorf("SonoAnagrammi(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestFirmaRune(t *testing.T) {
	f1 := FirmaRune("listen")
	f2 := FirmaRune("silent")
	if len(f1) != len(f2) {
		t.Fatalf("FirmaRune: lunghezze diverse: %d vs %d", len(f1), len(f2))
	}
	for i := range f1 {
		if f1[i] != f2[i] {
			t.Errorf("FirmaRune: firme diverse per anagrammi: %v vs %v", f1, f2)
			break
		}
	}
}
