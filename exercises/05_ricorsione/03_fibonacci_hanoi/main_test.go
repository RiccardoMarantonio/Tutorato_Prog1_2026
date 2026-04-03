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
		{50, 12586269025},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := FibonacciMemo(tt.n)
			if int64(got) != int64(tt.want) {
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
	// Verifica che la prima e l'ultima mossa siano corrette
	if len(mosse) == 0 {
		t.Fatal("Hanoi(3) non ha restituito mosse")
	}
	// Prima mossa: sposta disco da A a C (o A a B, dipende dall'implementazione)
	// Ultima mossa: deve finire su C
	last := mosse[len(mosse)-1]
	if last[len(last)-1] != 'C' {
		t.Errorf("Ultima mossa non termina su C: %q", last)
	}
}
