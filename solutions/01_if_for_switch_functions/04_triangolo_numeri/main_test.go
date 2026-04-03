package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCalcolaNumeroRiga(t *testing.T) {
	tests := []struct {
		riga int
		want int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CalcolaNumeroRiga(tt.riga); got != tt.want {
				t.Errorf("CalcolaNumeroRiga(%d) = %d, want %d", tt.riga, got, tt.want)
			}
		})
	}
}

func TestSommaTriangolo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 6},
		{3, 21},
		{4, 55},
		{5, 120},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SommaTriangolo(tt.n); got != tt.want {
				t.Errorf("SommaTriangolo(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestStampaTriangolo(t *testing.T) {
	capture := func(fn func()) string {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		fn()
		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		io.Copy(&buf, r)
		r.Close()
		return buf.String()
	}

	out := capture(func() { StampaTriangolo(3) })
	if out != "   1\n   2   3\n   4   5   6\n" {
		t.Errorf("StampaTriangolo(3) output non atteso:\n%q", out)
	}
}
