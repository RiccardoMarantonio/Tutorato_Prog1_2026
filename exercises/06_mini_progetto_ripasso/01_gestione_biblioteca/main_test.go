package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runExercise(input string) (string, error) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func TestGestioneBiblioteca(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "flusso completo",
			input: "AGGIUNGI 978-001 Il Nome Della Rosa Umberto Eco 1980 Narrativa\n" +
				"AGGIUNGI 978-002 1984 George Orwell 1949 Narrativa\n" +
				"AGGIUNGI 978-003 Clean Code Robert Martin 2008 Informatica\n" +
				"PRESTITO 978-001 Mario 15/01/2025\n" +
				"PER_GENERI\n" +
				"PRESTITI_APERTI\n" +
				"STATISTICHE\n" +
				"FINE",
			contains: []string{
				"Libro aggiunto: Il Nome Della Rosa",
				"Libro aggiunto: 1984",
				"Libro aggiunto: Clean Code",
				"Prestito registrato: 978-001 a Mario",
				"Libri per genere",
				"Narrativa: 2",
				"Informatica: 1",
				"Prestiti aperti",
				"978-001",
				"Mario",
				"Statistiche",
				"Totale libri: 3",
				"Disponibili: 2",
				"In prestito: 1",
				"Generi diversi: 2",
			},
		},
		{
			name: "cerca libro",
			input: "AGGIUNGI 978-001 Il Nome Della Rosa Umberto Eco 1980 Narrativa\n" +
				"AGGIUNGI 978-002 1984 George Orwell 1949 Narrativa\n" +
				"CERCA rosa\n" +
				"FINE",
			contains: []string{
				"Risultati ricerca",
				"Il Nome Della Rosa",
			},
		},
		{
			name: "cerca per autore",
			input: "AGGIUNGI 978-002 1984 George Orwell 1949 Narrativa\n" +
				"CERCA orwell\n" +
				"FINE",
			contains: []string{
				"1984",
			},
		},
		{
			name: "restituzione",
			input: "AGGIUNGI 978-001 Il Nome Della Rosa Umberto Eco 1980 Narrativa\n" +
				"PRESTITO 978-001 Mario 15/01/2025\n" +
				"RESTITUZIONE 978-001 20/02/2025\n" +
				"PRESTITI_APERTI\n" +
				"STATISTICHE\n" +
				"FINE",
			contains: []string{
				"Restituzione registrata: 978-001",
				"Disponibili: 1",
				"In prestito: 0",
			},
		},
		{
			name: "catalogo",
			input: "AGGIUNGI 978-001 Il Nome Della Rosa Umberto Eco 1980 Narrativa\n" +
				"CATALOGO\n" +
				"FINE",
			contains: []string{
				"Catalogo completo",
				"DISPONIBILE",
			},
		},
		{
			name: "prestito libro non disponibile",
			input: "AGGIUNGI 978-001 Il Nome Della Rosa Umberto Eco 1980 Narrativa\n" +
				"PRESTITO 978-001 Mario 15/01/2025\n" +
				"PRESTITO 978-001 Luigi 20/01/2025\n" +
				"FINE",
			contains: []string{
				"Errore",
			},
		},
		{
			name: "duplicato isbn",
			input: "AGGIUNGI 978-001 Il Nome Della Rosa Umberto Eco 1980 Narrativa\n" +
				"AGGIUNGI 978-001 Altro Libro Autore 2000 Genere\n" +
				"FINE",
			contains: []string{
				"Errore",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			for _, c := range tt.contains {
				if !strings.Contains(got, c) {
					t.Errorf("output non contiene %q:\n%s", c, got)
				}
			}
		})
	}
}
