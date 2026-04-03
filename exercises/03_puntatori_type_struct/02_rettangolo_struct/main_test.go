package main

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		name string
		r    Rettangolo
		want float64
	}{
		{"base", Rettangolo{4, 6}, 24},
		{"quadrato", Rettangolo{5, 5}, 25},
		{"zero", Rettangolo{0, 10}, 0},
		{"decimale", Rettangolo{2.5, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Area(tt.r)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 0.01 {
				t.Errorf("Area(%v) = %.2f, want %.2f", tt.r, got, tt.want)
			}
		})
	}
}

func TestPerimetro(t *testing.T) {
	tests := []struct {
		name string
		r    Rettangolo
		want float64
	}{
		{"base", Rettangolo{4, 6}, 20},
		{"quadrato", Rettangolo{5, 5}, 20},
		{"zero", Rettangolo{0, 10}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Perimetro(tt.r)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 0.01 {
				t.Errorf("Perimetro(%v) = %.2f, want %.2f", tt.r, got, tt.want)
			}
		})
	}
}

func TestScala(t *testing.T) {
	r := Rettangolo{4, 6}
	Scala(&r, 2)
	if r.Base != 8 || r.Altezza != 12 {
		t.Errorf("Scala x2: Base = %.2f, Altezza = %.2f, want 8, 12", r.Base, r.Altezza)
	}
}

func TestScalaNonModificaOriginale(t *testing.T) {
	r := Rettangolo{3, 5}
	copia := r
	Scala(&r, 3)
	if copia.Base != 3 || copia.Altezza != 5 {
		t.Errorf("La copia è stata modificata: Base = %.2f, Altezza = %.2f", copia.Base, copia.Altezza)
	}
}
