package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{10, 55},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFibonacciFinoA(t *testing.T) {
	got := FibonacciFinoA(10)
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if len(got) != len(want) {
		t.Errorf("FibonacciFinoA(10) length = %d, want %d", len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("FibonacciFinoA(10)[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}
