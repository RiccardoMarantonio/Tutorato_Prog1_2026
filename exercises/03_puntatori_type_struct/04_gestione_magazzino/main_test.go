package main

import "testing"

func TestAggiungiProdotto(t *testing.T) {
	m := &Magazzino{}
	err := AggiungiProdotto(m, Prodotto{"P001", "Laptop", 899.99, 10})
	if err != nil {
		t.Errorf("AggiungiProdotto: errore inatteso: %v", err)
	}

	err2 := AggiungiProdotto(m, Prodotto{"P001", "Laptop2", 500, 5})
	if err2 == nil {
		t.Error("AggiungiProdotto duplicato: errore atteso, got nil")
	}

	if len(m.Prodotti) != 1 {
		t.Errorf("len(Prodotti) = %d, want 1", len(m.Prodotti))
	}
}

func TestRegistraCarico(t *testing.T) {
	m := &Magazzino{}
	AggiungiProdotto(m, Prodotto{"P001", "Laptop", 899.99, 10})

	err := RegistraCarico(m, "P001", 5)
	if err != nil {
		t.Errorf("RegistraCarico: errore inatteso: %v", err)
	}
	if m.Prodotti[0].Quantita != 15 {
		t.Errorf("Quantita dopo carico = %d, want 15", m.Prodotti[0].Quantita)
	}

	err = RegistraCarico(m, "P999", 1)
	if err == nil {
		t.Error("RegistraCarico prodotto inesistente: errore atteso, got nil")
	}

	err = RegistraCarico(m, "P001", 0)
	if err == nil {
		t.Error("RegistraCarico quantita zero: errore atteso, got nil")
	}
}

func TestRegistraScarico(t *testing.T) {
	m := &Magazzino{}
	AggiungiProdotto(m, Prodotto{"P001", "Laptop", 899.99, 10})

	err := RegistraScarico(m, "P001", 3)
	if err != nil {
		t.Errorf("RegistraScarico: errore inatteso: %v", err)
	}
	if m.Prodotti[0].Quantita != 7 {
		t.Errorf("Quantita dopo scarico = %d, want 7", m.Prodotti[0].Quantita)
	}

	err = RegistraScarico(m, "P001", 100)
	if err == nil {
		t.Error("RegistraScarico quantita eccessiva: errore atteso, got nil")
	}

	err = RegistraScarico(m, "P999", 1)
	if err == nil {
		t.Error("RegistraScarico prodotto inesistente: errore atteso, got nil")
	}
}

func TestValoreMagazzino(t *testing.T) {
	m := &Magazzino{Prodotti: []Prodotto{
		{"P001", "A", 10.0, 5},
		{"P002", "B", 20.0, 3},
	}}
	got := ValoreMagazzino(*m)
	want := 110.0
	diff := got - want
	if diff < 0 {
		diff = -diff
	}
	if diff > 0.01 {
		t.Errorf("ValoreMagazzino = %.2f, want %.2f", got, want)
	}
}

func TestProdottiSottoScorta(t *testing.T) {
	m := &Magazzino{Prodotti: []Prodotto{
		{"P001", "A", 10.0, 5},
		{"P002", "B", 20.0, 15},
		{"P003", "C", 5.0, 3},
	}}
	count := ProdottiSottoScorta(*m, 10)
	if count != 2 {
		t.Errorf("ProdottiSottoScorta(10) = %d, want 2", count)
	}
}

func TestCercaProdotto(t *testing.T) {
	m := &Magazzino{Prodotti: []Prodotto{
		{"P001", "Laptop", 899.99, 10},
	}}

	p, err := CercaProdotto(*m, "P001")
	if err != nil {
		t.Errorf("CercaProdotto: errore inatteso: %v", err)
	}
	if p.Nome != "Laptop" {
		t.Errorf("CercaProdotto: Nome = %q, want %q", p.Nome, "Laptop")
	}

	_, err2 := CercaProdotto(*m, "P999")
	if err2 == nil {
		t.Error("CercaProdotto inesistente: errore atteso, got nil")
	}
}
