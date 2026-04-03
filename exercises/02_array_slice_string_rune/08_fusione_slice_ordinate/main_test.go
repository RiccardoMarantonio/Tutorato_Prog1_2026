package main

import "testing"

func TestMergeSorted(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"base", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"vuota a", []int{}, []int{1, 2}, []int{1, 2}},
		{"vuota b", []int{1, 2}, []int{}, []int{1, 2}},
		{"entrambe vuote", []int{}, []int{}, []int{}},
		{"duplicati", []int{1, 3}, []int{1, 3}, []int{1, 1, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSorted(tt.a, tt.b)
			if len(got) != len(tt.want) {
				t.Errorf("MergeSorted(%v, %v) length = %d, want %d", tt.a, tt.b, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("MergeSorted(%v, %v)[%d] = %d, want %d", tt.a, tt.b, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestRimuoviDuplicati(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"con duplicati", []int{1, 2, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"senza duplicati", []int{1, 2, 3}, []int{1, 2, 3}},
		{"vuoto", []int{}, []int{}},
		{"tutti uguali", []int{5, 5, 5}, []int{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RimuoviDuplicati(tt.slice)
			if len(got) != len(tt.want) {
				t.Errorf("RimuoviDuplicati(%v) length = %d, want %d", tt.slice, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("RimuoviDuplicati(%v)[%d] = %d, want %d", tt.slice, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestIntersezione(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"base", []int{1, 3, 5, 7}, []int{3, 5, 9}, []int{3, 5}},
		{"nessuno", []int{1, 2}, []int{3, 4}, []int{}},
		{"vuoto", []int{}, []int{1, 2}, []int{}},
		{"identiche", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersezione(tt.a, tt.b)
			if len(got) != len(tt.want) {
				t.Errorf("Intersezione(%v, %v) length = %d, want %d", tt.a, tt.b, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Intersezione(%v, %v)[%d] = %d, want %d", tt.a, tt.b, i, got[i], tt.want[i])
				}
			}
		})
	}
}
