package main

import "testing"

func TestHanoi(t *testing.T) {
	tests := []struct {
		n     int
		mosse int
	}{
		{1, 1},
		{2, 3},
		{3, 7},
		{4, 15},
		{5, 31},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			mosse := Hanoi(tt.n, "A", "B", "C")
			if len(mosse) != tt.mosse {
				t.Errorf("Hanoi(%d) = %d mosse, want %d", tt.n, len(mosse), tt.mosse)
			}
		})
	}
}

func TestHanoiMosseValide(t *testing.T) {
	mosse := Hanoi(3, "A", "B", "C")
	if len(mosse) == 0 {
		t.Fatal("Hanoi(3) non ha restituito mosse")
	}
	last := mosse[len(mosse)-1]
	if last[len(last)-1] != 'C' {
		t.Errorf("Ultima mossa non termina su C: %q", last)
	}
}
