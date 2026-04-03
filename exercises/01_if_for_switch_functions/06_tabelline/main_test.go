package main

import "testing"

func TestStampaTabellina(t *testing.T) {
	tests := []struct {
		n     int
		linee int
	}{
		{1, 11}, // 10 righe + 1 riga vuota
		{2, 22}, // 2 * (10 + 1)
		{3, 33}, // 3 * (10 + 1)
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			out := catturaOutput(func() { StampaTabelline(tt.n) })
			righe := contaRighe(out)
			if righe != tt.linee {
				t.Errorf("StampaTabelline(%d) = %d righe, want %d", tt.n, righe, tt.linee)
			}
		})
	}
}

func TestRigaTabellina(t *testing.T) {
	tests := []struct {
		base int
		molt int
		want string
	}{
		{1, 1, "1 x 1 = 1"},
		{2, 3, "2 x 3 = 6"},
		{7, 8, "7 x 8 = 56"},
		{12, 10, "12 x 10 = 120"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RigaTabellina(tt.base, tt.molt)
			if got != tt.want {
				t.Errorf("RigaTabellina(%d, %d) = %q, want %q", tt.base, tt.molt, got, tt.want)
			}
		})
	}
}
