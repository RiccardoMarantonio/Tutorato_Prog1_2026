package main

import "testing"

func TestAggiungiVoto(t *testing.T) {
	registro := make(map[string]*Studente)

	err := AggiungiVoto(registro, "Mario", 25)
	if err != nil {
		t.Errorf("AggiungiVoto: errore inatteso: %v", err)
	}
	if registro["Mario"] == nil {
		t.Fatal("Studente Mario non creato")
	}
	if len(registro["Mario"].Voti) != 1 || registro["Mario"].Voti[0] != 25 {
		t.Errorf("Voti Mario = %v, want [25]", registro["Mario"].Voti)
	}

	AggiungiVoto(registro, "Mario", 28)
	if len(registro["Mario"].Voti) != 2 {
		t.Errorf("Mario ha %d voti, want 2", len(registro["Mario"].Voti))
	}

	err = AggiungiVoto(registro, "Luigi", 31)
	if err == nil {
		t.Error("AggiungiVoto(31): errore atteso, got nil")
	}
	err = AggiungiVoto(registro, "Luigi", 0)
	if err == nil {
		t.Error("AggiungiVoto(0): errore atteso, got nil")
	}
}

func TestMediaStudente(t *testing.T) {
	registro := make(map[string]*Studente)
	AggiungiVoto(registro, "Mario", 25)
	AggiungiVoto(registro, "Mario", 28)
	AggiungiVoto(registro, "Mario", 24)

	media, err := MediaStudente(registro, "Mario")
	if err != nil {
		t.Fatalf("MediaStudente: errore inatteso: %v", err)
	}
	diff := media - 25.67
	if diff < 0 {
		diff = -diff
	}
	if diff > 0.01 {
		t.Errorf("MediaStudente(Mario) = %.2f, want 25.67", media)
	}

	_, err = MediaStudente(registro, "Inesistente")
	if err == nil {
		t.Error("MediaStudente inesistente: errore atteso, got nil")
	}
}

func TestMigliorPeggiore(t *testing.T) {
	registro := make(map[string]*Studente)
	AggiungiVoto(registro, "Mario", 25)
	AggiungiVoto(registro, "Mario", 28)
	AggiungiVoto(registro, "Luigi", 18)
	AggiungiVoto(registro, "Luigi", 22)
	AggiungiVoto(registro, "Anna", 30)

	migliore, peggiore, err := MigliorPeggiore(registro)
	if err != nil {
		t.Fatalf("MigliorPeggiore: errore inatteso: %v", err)
	}
	if migliore != "Anna" {
		t.Errorf("Migliore = %q, want %q", migliore, "Anna")
	}
	if peggiore != "Luigi" {
		t.Errorf("Peggiore = %q, want %q", peggiore, "Luigi")
	}

	registro2 := make(map[string]*Studente)
	_, _, err = MigliorPeggiore(registro2)
	if err == nil {
		t.Error("MigliorPeggiore vuoto: errore atteso, got nil")
	}
}

func TestDistribuzioneVoti(t *testing.T) {
	registro := make(map[string]*Studente)
	AggiungiVoto(registro, "A", 15)
	AggiungiVoto(registro, "B", 20)
	AggiungiVoto(registro, "C", 25)
	AggiungiVoto(registro, "D", 28)
	AggiungiVoto(registro, "E", 30)

	dist := DistribuzioneVoti(registro)
	if dist["Insufficiente"] != 1 {
		t.Errorf("Insufficiente: %d, want 1", dist["Insufficiente"])
	}
	if dist["Sufficiente"] != 1 {
		t.Errorf("Sufficiente: %d, want 1", dist["Sufficiente"])
	}
	if dist["Buono"] != 1 {
		t.Errorf("Buono: %d, want 1", dist["Buono"])
	}
	if dist["Ottimo"] != 2 {
		t.Errorf("Ottimo: %d, want 2", dist["Ottimo"])
	}
}
