package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestLeggiFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")

	content := "ciao mondo ciao"
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("errore creazione file: %v", err)
	}

	luogo, err := LeggiFile(testFile)
	if err != nil {
		t.Fatalf("LeggiFile fallita: %v", err)
	}
	if luogo != content {
		t.Errorf("contenuto non corrisponde: got %q, want %q", luogo, content)
	}
}

func TestEstraiParole(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"ciao mondo", []string{"ciao", "mondo"}},
		{"Ciao MONDO", []string{"ciao", "mondo"}},
		{"ciao, mondo!", []string{"ciao", "mondo"}},
	}

	for _, tt := range tests {
		got := EstraiParole(tt.input)
		if len(got) != len(tt.expected) {
			t.Errorf("EstraiParole(%q) lunghezza errata: got %d, want %d", tt.input, len(got), len(tt.expected))
			continue
		}
		for i, w := range got {
			if w != tt.expected[i] {
				t.Errorf("EstraiParole(%q)[%d]: got %q, want %q", tt.input, i, w, tt.expected[i])
			}
		}
	}
}

func TestContaFrequenze(t *testing.T) {
	parole := []string{"ciao", "mondo", "ciao", "ciao"}
	freq := ContaFrequenze(parole)

	if freq["ciao"] != 3 {
		t.Errorf("freq[ciao]: got %d, want 3", freq["ciao"])
	}
	if freq["mondo"] != 1 {
		t.Errorf("freq[mondo]: got %d, want 1", freq["mondo"])
	}
}

func TestOrdinaPerFrequenza(t *testing.T) {
	freq := map[string]int{"ciao": 3, "mondo": 1, "test": 2}
	ordinato := OrdinaPerFrequenza(freq)

	if len(ordinato) != 3 {
		t.Fatalf("lunghezza errata: got %d, want 3", len(ordinato))
	}
	if ordinato[0].Parola != "ciao" || ordinato[0].Frequenza != 3 {
		t.Errorf("primo elemento errato: %+v", ordinato[0])
	}
	if ordinato[1].Parola != "test" || ordinato[1].Frequenza != 2 {
		t.Errorf("secondo elemento errato: %+v", ordinato[1])
	}
	if ordinato[2].Parola != "mondo" || ordinato[2].Frequenza != 1 {
		t.Errorf("terzo elemento errato: %+v", ordinato[2])
	}
}

func runExercise(filePath string) (string, error) {
	cmd := exec.Command("go", "run", ".", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func TestIstogramma(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")

	content := "il gatto mangia il pesce il gatto"
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("errore creazione file: %v", err)
	}

	output, err := runExercise(testFile)
	if err != nil {
		t.Fatalf("esecuzione fallita: %v\n%s", err, output)
	}

	if !bytes.Contains([]byte(output), []byte("il")) {
		t.Errorf("output deve contenere 'il':\n%s", output)
	}
	if !bytes.Contains([]byte(output), []byte("gatto")) {
		t.Errorf("output deve contenere 'gatto':\n%s", output)
	}
}
