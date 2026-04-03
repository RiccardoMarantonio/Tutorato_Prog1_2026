package main

import "testing"

func TestTrovaMax(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"base", []int{10, 3, 7, 1, 9}, 10},
		{"negativi", []int{-5, -2, -9, -1}, -1},
		{"singolo", []int{42}, 42},
		{"vuoto", []int{}, 0},
		{"duplicati", []int{5, 5, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrovaMax(tt.slice); got != tt.want {
				t.Errorf("TrovaMax(%v) = %d, want %d", tt.slice, got, tt.want)
			}
		})
	}
}

func TestTrovaMin(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"base", []int{10, 3, 7, 1, 9}, 1},
		{"negativi", []int{-5, -2, -9, -1}, -9},
		{"singolo", []int{42}, 42},
		{"vuoto", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrovaMin(tt.slice); got != tt.want {
				t.Errorf("TrovaMin(%v) = %d, want %d", tt.slice, got, tt.want)
			}
		})
	}
}

func TestTrovaMaxMin(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		wantMax int
		wantMin int
	}{
		{"base", []int{10, 3, 7, 1, 9}, 10, 1},
		{"vuoto", []int{}, 0, 0},
		{"singolo", []int{7}, 7, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotMin := TrovaMaxMin(tt.slice)
			if gotMax != tt.wantMax || gotMin != tt.wantMin {
				t.Errorf("TrovaMaxMin(%v) = (%d, %d), want (%d, %d)", tt.slice, gotMax, gotMin, tt.wantMax, tt.wantMin)
			}
		})
	}
}
