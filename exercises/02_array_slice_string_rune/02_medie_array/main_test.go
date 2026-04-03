package main

import "testing"

func TestMediaArray(t *testing.T) {
	tests := []struct {
		name string
		arr  [10]int
		want float64
	}{
		{"base", [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
		{"tutti uguali", [10]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, 5.0},
		{"negativi", [10]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4}, -0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MediaArray(tt.arr)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 0.01 {
				t.Errorf("MediaArray(%v) = %.2f, want %.2f", tt.arr, got, tt.want)
			}
		})
	}
}

func TestContaSopraMedia(t *testing.T) {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	media := 5.5
	sopra := ContaSopraMedia(arr, media)
	if sopra != 5 {
		t.Errorf("ContaSopraMedia = %d, want 5", sopra)
	}
}

func TestContaSottoMedia(t *testing.T) {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	media := 5.5
	sotto := ContaSottoMedia(arr, media)
	if sotto != 5 {
		t.Errorf("ContaSottoMedia = %d, want 5", sotto)
	}
}
