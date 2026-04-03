package main

import "testing"

func newTestElezione() *Elezione {
	return &Elezione{
		Candidati: []string{"Rossi", "Bianchi", "Verdi"},
		Sezioni: map[string]*Sezione{
			"Centro": {Nome: "Centro", Voti: make(map[string]int), Elettori: 100},
			"Nord":   {Nome: "Nord", Voti: make(map[string]int), Elettori: 80},
		},
	}
}

func TestRegistraVoto(t *testing.T) {
	e := newTestElezione()

	err := RegistraVoto(e, "Centro", "Rossi")
	if err != nil {
		t.Errorf("RegistraVoto: errore inatteso: %v", err)
	}
	if e.Sezioni["Centro"].Voti["Rossi"] != 1 {
		t.Errorf("Voti Centro[Rossi] = %d, want 1", e.Sezioni["Centro"].Voti["Rossi"])
	}
	if e.Sezioni["Centro"].Votanti != 1 {
		t.Errorf("Votanti Centro = %d, want 1", e.Sezioni["Centro"].Votanti)
	}

	err = RegistraVoto(e, "Sud", "Rossi")
	if err == nil {
		t.Error("RegistraVoto sezione inesistente: errore atteso, got nil")
	}

	err = RegistraVoto(e, "Centro", "Neri")
	if err == nil {
		t.Error("RegistraVoto candidato non valido: errore atteso, got nil")
	}
}

func TestRisultatiNazionali(t *testing.T) {
	e := newTestElezione()
	RegistraVoto(e, "Centro", "Rossi")
	RegistraVoto(e, "Centro", "Rossi")
	RegistraVoto(e, "Nord", "Bianchi")
	RegistraVoto(e, "Nord", "Verdi")

	ris := RisultatiNazionali(*e)
	if ris["Rossi"] != 2 {
		t.Errorf("Rossi: %d, want 2", ris["Rossi"])
	}
	if ris["Bianchi"] != 1 {
		t.Errorf("Bianchi: %d, want 1", ris["Bianchi"])
	}
	if ris["Verdi"] != 1 {
		t.Errorf("Verdi: %d, want 1", ris["Verdi"])
	}
}

func TestAffluenza(t *testing.T) {
	e := newTestElezione()
	RegistraVoto(e, "Centro", "Rossi")
	RegistraVoto(e, "Centro", "Bianchi")

	aff := Affluenza(*e)
	diff := aff["Centro"] - 2.0
	if diff < 0 {
		diff = -diff
	}
	if diff > 0.01 {
		t.Errorf("Affluenza Centro = %.2f%%, want 2.00%%", aff["Centro"])
	}
	if aff["Nord"] != 0.0 {
		t.Errorf("Affluenza Nord = %.2f%%, want 0.00%%", aff["Nord"])
	}
}

func TestVincitore(t *testing.T) {
	e := newTestElezione()
	RegistraVoto(e, "Centro", "Rossi")
	RegistraVoto(e, "Centro", "Rossi")
	RegistraVoto(e, "Nord", "Bianchi")

	vinc, err := Vincitore(*e)
	if err != nil {
		t.Fatalf("Vincitore: errore inatteso: %v", err)
	}
	if vinc != "Rossi" {
		t.Errorf("Vincitore = %q, want %q", vinc, "Rossi")
	}

	e2 := newTestElezione()
	_, err = Vincitore(*e2)
	if err == nil {
		t.Error("Vincitore senza voti: errore atteso, got nil")
	}
}

func TestBallottaggio(t *testing.T) {
	e := newTestElezione()
	RegistraVoto(e, "Centro", "Rossi")
	RegistraVoto(e, "Centro", "Rossi")
	RegistraVoto(e, "Nord", "Bianchi")
	RegistraVoto(e, "Nord", "Bianchi")
	RegistraVoto(e, "Centro", "Verdi")

	c1, c2, err := Ballottaggio(*e)
	if err != nil {
		t.Fatalf("Ballottaggio: errore inatteso: %v", err)
	}
	if (c1 != "Rossi" && c1 != "Bianchi") || (c2 != "Rossi" && c2 != "Bianchi") {
		t.Errorf("Ballottaggio = %s vs %s, want Rossi vs Bianchi", c1, c2)
	}
}
