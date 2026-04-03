package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"zero", 0, true},
		{"due", 2, true},
		{"tre", 3, false},
		{"uno", 1, false},
		{"cento", 100, true},
		{"novantanove", 99, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.n); got != tt.want {
				t.Errorf("IsEven(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
