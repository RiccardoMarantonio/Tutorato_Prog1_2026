package main

import (
	"fmt"
	"testing"
)

func TestIsMultiplo(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		divisore int
		want     bool
	}{
		{"multiplo di 3", 9, 3, true},
		{"non multiplo di 3", 10, 3, false},
		{"multiplo di 5", 20, 5, true},
		{"multiplo di se stesso", 7, 7, true},
		{"zero", 0, 5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMultiplo(tt.n, tt.divisore); got != tt.want {
				t.Errorf("IsMultiplo(%d, %d) = %v, want %v", tt.n, tt.divisore, got, tt.want)
			}
		})
	}
}

func TestIsPrimo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"uno non è primo", 1, false},
		{"due è primo", 2, true},
		{"tre è primo", 3, true},
		{"quattro non è primo", 4, false},
		{"undici è primo", 11, true},
		{"quindici non è primo", 15, false},
		{"ventinove è primo", 29, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimo(tt.n); got != tt.want {
				t.Errorf("IsPrimo(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{2, "Prime"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "Prime"},
		{12, "Fizz"},
		{13, "Prime"},
		{15, "FizzBuzz"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("FizzBuzz(%d)", tt.n), func(t *testing.T) {
			if got := FizzBuzz(tt.n); got != tt.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}
