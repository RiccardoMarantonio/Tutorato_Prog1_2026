package main

import "testing"

func TestAggiungi(t *testing.T) {
	r := &Rubrica{}
	Aggiungi(r, Contatto{"Mario", "Rossi", "333111"})
	Aggiungi(r, Contatto{"Luigi", "Bianchi", "333222"})

	if len(r.Contatti) != 2 {
		t.Errorf("len(Contatti) = %d, want 2", len(r.Contatti))
	}

	Aggiungi(r, Contatto{"Mario", "Verdi", "333111"})
	if len(r.Contatti) != 2 {
		t.Errorf("duplicato aggiunto: len(Contatti) = %d, want 2", len(r.Contatti))
	}
}

func TestCercaPerNome(t *testing.T) {
	r := Rubrica{Contatti: []Contatto{
		{"Mario", "Rossi", "333111"},
		{"Luigi", "Bianchi", "333222"},
		{"Mario", "Verdi", "333333"},
	}}

	results := CercaPerNome(r, "Mario")
	if len(results) != 2 {
		t.Errorf("CercaPerNome(Mario) = %d risultati, want 2", len(results))
	}

	results2 := CercaPerNome(r, "mario")
	if len(results2) != 2 {
		t.Errorf("CercaPerNome(mario) = %d risultati, want 2", len(results2))
	}

	results3 := CercaPerNome(r, "Luigi")
	if len(results3) != 1 {
		t.Errorf("CercaPerNome(Luigi) = %d risultati, want 1", len(results3))
	}
}

func TestRimuovi(t *testing.T) {
	r := &Rubrica{Contatti: []Contatto{
		{"Mario", "Rossi", "333111"},
		{"Luigi", "Bianchi", "333222"},
	}}

	if !Rimuovi(r, "333111") {
		t.Error("Rimuovi(333111) = false, want true")
	}
	if len(r.Contatti) != 1 {
		t.Errorf("len(Contatti) dopo rimozione = %d, want 1", len(r.Contatti))
	}
	if r.Contatti[0].Nome != "Luigi" {
		t.Errorf("Contatto rimasto: %s, want Luigi", r.Contatti[0].Nome)
	}

	if Rimuovi(r, "000000") {
		t.Error("Rimuovi(000000) = true, want false")
	}
}
