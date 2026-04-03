package main

import "testing"

func TestFibonacciMemo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := FibonacciMemo(tt.n)
			if got != tt.want {
				t.Errorf("FibonacciMemo(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestContaChiamate(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 3},
		{3, 5},
		{4, 9},
		{5, 15},
		{10, 177},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ContaChiamate(tt.n); got != tt.want {
				t.Errorf("ContaChiamate(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
