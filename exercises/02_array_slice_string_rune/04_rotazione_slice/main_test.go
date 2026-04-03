package main

import "testing"

func TestRuota(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		k     int
		want  []int
	}{
		{"base", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"k zero", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"k maggiore", []int{1, 2, 3}, 5, []int{2, 3, 1}},
		{"vuoto", []int{}, 3, []int{}},
		{"singolo", []int{7}, 10, []int{7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ruota(tt.slice, tt.k)
			if len(got) != len(tt.want) {
				t.Errorf("Ruota(%v, %d) length = %d, want %d", tt.slice, tt.k, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Ruota(%v, %d)[%d] = %d, want %d", tt.slice, tt.k, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestRuotaNonModificaOriginale(t *testing.T) {
	orig := []int{1, 2, 3}
	Ruota(orig, 1)
	if orig[0] != 1 || orig[1] != 2 || orig[2] != 3 {
		t.Errorf("Ruota ha modificato la slice originale: %v", orig)
	}
}
