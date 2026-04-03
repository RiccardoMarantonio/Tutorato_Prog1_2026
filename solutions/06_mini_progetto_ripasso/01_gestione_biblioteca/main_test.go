package main

import "testing"

func newTestBiblioteca() *Biblioteca {
	b := NuovaBiblioteca()
	AggiungiLibro(b, Libro{"978-001", "Il Nome Della Rosa", "Umberto Eco", 1980, "Narrativa", true})
	AggiungiLibro(b, Libro{"978-002", "1984", "George Orwell", 1949, "Narrativa", true})
	AggiungiLibro(b, Libro{"978-003", "Clean Code", "Robert Martin", 2008, "Informatica", true})
	return b
}

func TestNuovaBiblioteca(t *testing.T) {
	b := NuovaBiblioteca()
	if b == nil {
		t.Fatal("NuovaBiblioteca() = nil")
	}
	if b.Libri == nil {
		t.Error("b.Libri = nil")
	}
	if b.Prestiti == nil {
		t.Error("b.Prestiti = nil")
	}
}

func TestAggiungiLibro(t *testing.T) {
	b := NuovaBiblioteca()
	err := AggiungiLibro(b, Libro{"978-001", "Test", "Autore", 2020, "Genere", true})
	if err != nil {
		t.Errorf("AggiungiLibro: errore inatteso: %v", err)
	}
	if len(b.Libri) != 1 {
		t.Errorf("len(Libri) = %d, want 1", len(b.Libri))
	}

	err2 := AggiungiLibro(b, Libro{"978-001", "Test2", "Autore2", 2021, "Genere2", true})
	if err2 == nil {
		t.Error("AggiungiLibro duplicato: errore atteso, got nil")
	}
}

func TestCercaLibro(t *testing.T) {
	b := newTestBiblioteca()

	results := CercaLibro(*b, "rosa")
	if len(results) != 1 {
		t.Errorf("CercaLibro(rosa) = %d risultati, want 1", len(results))
	}

	results2 := CercaLibro(*b, "orwell")
	if len(results2) != 1 || results2[0].Titolo != "1984" {
		t.Errorf("CercaLibro(orwell) = %v, want [1984]", results2)
	}

	results3 := CercaLibro(*b, "inesistente")
	if len(results3) != 0 {
		t.Errorf("CercaLibro(inesistente) = %d risultati, want 0", len(results3))
	}
}

func TestRegistraPrestito(t *testing.T) {
	b := newTestBiblioteca()

	err := RegistraPrestito(b, "978-001", "Mario", "15/01/2025")
	if err != nil {
		t.Errorf("RegistraPrestito: errore inatteso: %v", err)
	}
	if b.Libri["978-001"].Disponibile {
		t.Error("Libro dovrebbe essere non disponibile dopo prestito")
	}
	if len(b.Prestiti) != 1 {
		t.Errorf("len(Prestiti) = %d, want 1", len(b.Prestiti))
	}

	err2 := RegistraPrestito(b, "978-001", "Luigi", "20/01/2025")
	if err2 == nil {
		t.Error("RegistraPrestito libro gia in prestito: errore atteso, got nil")
	}

	err3 := RegistraPrestito(b, "978-999", "Anna", "25/01/2025")
	if err3 == nil {
		t.Error("RegistraPrestito libro inesistente: errore atteso, got nil")
	}
}

func TestRegistraRestituzione(t *testing.T) {
	b := newTestBiblioteca()
	RegistraPrestito(b, "978-001", "Mario", "15/01/2025")

	err := RegistraRestituzione(b, "978-001", "20/02/2025")
	if err != nil {
		t.Errorf("RegistraRestituzione: errore inatteso: %v", err)
	}
	if !b.Libri["978-001"].Disponibile {
		t.Error("Libro dovrebbe essere disponibile dopo restituzione")
	}

	err2 := RegistraRestituzione(b, "978-001", "21/02/2025")
	if err2 == nil {
		t.Error("RegistraRestituzione libro non in prestito: errore atteso, got nil")
	}
}

func TestLibriPerGenere(t *testing.T) {
	b := newTestBiblioteca()
	gen := LibriPerGenere(*b)
	if gen["Narrativa"] != 2 {
		t.Errorf("Narrativa: %d, want 2", gen["Narrativa"])
	}
	if gen["Informatica"] != 1 {
		t.Errorf("Informatica: %d, want 1", gen["Informatica"])
	}
}

func TestPrestitiAperti(t *testing.T) {
	b := newTestBiblioteca()
	RegistraPrestito(b, "978-001", "Mario", "15/01/2025")
	RegistraPrestito(b, "978-002", "Luigi", "20/01/2025")
	RegistraRestituzione(b, "978-002", "25/01/2025")

	aperti := PrestitiAperti(*b)
	if len(aperti) != 1 {
		t.Errorf("PrestitiAperti = %d, want 1", len(aperti))
	}
	if aperti[0].ISBN != "978-001" {
		t.Errorf("PrestitiAperti[0].ISBN = %q, want %q", aperti[0].ISBN, "978-001")
	}
}

func TestStatistiche(t *testing.T) {
	b := newTestBiblioteca()
	RegistraPrestito(b, "978-001", "Mario", "15/01/2025")

	stats := Statistiche(*b)
	if stats["totale"] != 3 {
		t.Errorf("totale: %d, want 3", stats["totale"])
	}
	if stats["disponibili"] != 2 {
		t.Errorf("disponibili: %d, want 2", stats["disponibili"])
	}
	if stats["in_prestito"] != 1 {
		t.Errorf("in_prestito: %d, want 1", stats["in_prestito"])
	}
	if stats["generi"] != 2 {
		t.Errorf("generi: %d, want 2", stats["generi"])
	}
}
