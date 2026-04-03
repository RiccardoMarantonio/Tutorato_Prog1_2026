package main

import "testing"

func TestScambia(t *testing.T) {
	a := 5
	b := 10
	Scambia(&a, &b)
	if a != 10 || b != 5 {
		t.Errorf("Scambia: a = %d, b = %d, want a = 10, b = 5", a, b)
	}
}

func TestScambiaUguali(t *testing.T) {
	a := 7
	b := 7
	Scambia(&a, &b)
	if a != 7 || b != 7 {
		t.Errorf("Scambia con valori uguali: a = %d, b = %d, want 7, 7", a, b)
	}
}

func TestScambiaNegativi(t *testing.T) {
	a := -3
	b := 8
	Scambia(&a, &b)
	if a != 8 || b != -3 {
		t.Errorf("Scambia negativi: a = %d, b = %d, want 8, -3", a, b)
	}
}
