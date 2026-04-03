package main

import "testing"

func TestAggiungiProdotto(t *testing.T) {
	m := &Magazzino{}
	err := AggiungiProdotto(m, Prodotto{"P001", "Laptop", "Elettronica", 899.99, 10})
	if err != nil {
		t.Errorf("AggiungiProdotto: errore inatteso: %v", err)
	}

	// Duplicato
	err2 := AggiungiProdotto(m, Prodotto{"P001", "Laptop2", "Elettronica", 500, 5})
	if err2 == nil {
		t.Error("AggiungiProdotto duplicato: errore atteso, got nil")
	}

	if len(m.Prodotti) != 1 {
		t.Errorf("len(Prodotti) = %d, want 1", len(m.Prodotti))
	}
}

func TestRegistraMovimento(t *testing.T) {
	m := &Magazzino{}
	AggiungiProdotto(m, Prodotto{"P001", "Laptop", "Elettronica", 899.99, 10})

	// Scarico valido
	err := RegistraMovimento(m, Movimento{"P001", "SCARICO", 3})
	if err != nil {
		t.Errorf("RegistraMovimento scarico: errore inatteso: %v", err)
	}
	if m.Prodotti[0].Quantita != 7 {
		t.Errorf("Quantita dopo scarico = %d, want 7", m.Prodotti[0].Quantita)
	}

	// Carico
	err = RegistraMovimento(m, Movimento{"P001", "CARICO", 5})
	if err != nil {
		t.Errorf("RegistraMovimento carico: errore inatteso: %v", err)
	}
	if m.Prodotti[0].Quantita != 12 {
		t.Errorf("Quantita dopo carico = %d, want 12", m.Prodotti[0].Quantita)
	}

	// Scarico troppo
	err = RegistraMovimento(m, Movimento{"P001", "SCARICO", 100})
	if err == nil {
		t.Error("RegistraMovimento scarico eccessivo: errore atteso, got nil")
	}

	// Prodotto inesistente
	err = RegistraMovimento(m, Movimento{"P999", "CARICO", 1})
	if err == nil {
		t.Error("RegistraMovimento prodotto inesistente: errore atteso, got nil")
	}

	// Quantita zero
	err = RegistraMovimento(m, Movimento{"P001", "CARICO", 0})
	if err == nil {
		t.Error("RegistraMovimento quantita zero: errore atteso, got nil")
	}
}

func TestValoreMagazzino(t *testing.T) {
	m := &Magazzino{Prodotti: []Prodotto{
		{"P001", "A", "Cat1", 10.0, 5},
		{"P002", "B", "Cat1", 20.0, 3},
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
		{"P001", "A", "Cat1", 10.0, 5},
		{"P002", "B", "Cat1", 20.0, 15},
		{"P003", "C", "Cat2", 5.0, 3},
	}}
	sotto := ProdottiSottoScorta(*m, 10)
	if len(sotto) != 2 {
		t.Errorf("ProdottiSottoScorta(10) = %d prodotti, want 2", len(sotto))
	}
}

func TestRiepilogoCategoria(t *testing.T) {
	m := &Magazzino{Prodotti: []Prodotto{
		{"P001", "A", "Elettronica", 10.0, 5},
		{"P002", "B", "Elettronica", 20.0, 3},
		{"P003", "C", "Arredo", 5.0, 10},
	}}
	riep := RiepilogoCategoria(*m)
	if riep["Elettronica"] != 8 {
		t.Errorf("Elettronica: %d, want 8", riep["Elettronica"])
	}
	if riep["Arredo"] != 10 {
		t.Errorf("Arredo: %d, want 10", riep["Arredo"])
	}
}
