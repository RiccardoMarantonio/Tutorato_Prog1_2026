package main

import "testing"

func TestSommaArray(t *testing.T) {
	tests := []struct {
		name string
		arr  [5]int
		want int
	}{
		{"tutti positivi", [5]int{1, 2, 3, 4, 5}, 15},
		{"tutti zeri", [5]int{0, 0, 0, 0, 0}, 0},
		{"con negativi", [5]int{-1, 2, -3, 4, -5}, -3},
		{"tutti uguali", [5]int{10, 10, 10, 10, 10}, 50},
		{"misti", [5]int{100, 0, -100, 50, 50}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SommaArray(tt.arr); got != tt.want {
				t.Errorf("SommaArray(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}
