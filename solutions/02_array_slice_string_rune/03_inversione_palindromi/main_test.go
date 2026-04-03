package main

import "testing"

func TestInvertiSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"base", []int{1, 2, 3}, []int{3, 2, 1}},
		{"vuoto", []int{}, []int{}},
		{"singolo", []int{5}, []int{5}},
		{"due", []int{10, 20}, []int{20, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvertiSlice(tt.slice)
			if len(got) != len(tt.want) {
				t.Errorf("InvertiSlice(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("InvertiSlice(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
			orig := []int{1, 2, 3}
			InvertiSlice(orig)
			for i := range orig {
				if orig[i] != []int{1, 2, 3}[i] {
					t.Errorf("InvertiSlice ha modificato la slice originale")
					break
				}
			}
		})
	}
}

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
		{"frase", "i topi non avevano nipoti", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EPalindromo(tt.s); got != tt.want {
				t.Errorf("EPalindromo(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestEPalindromoNumerico(t *testing.T) {
	tests := []struct {
		name   string
		numeri []int
		want   bool
	}{
		{"base", []int{1, 2, 3, 2, 1}, true},
		{"no", []int{1, 2, 3}, false},
		{"vuoto", []int{}, true},
		{"uguale", []int{5, 5, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EPalindromoNumerico(tt.numeri); got != tt.want {
				t.Errorf("EPalindromoNumerico(%v) = %v, want %v", tt.numeri, got, tt.want)
			}
		})
	}
}
