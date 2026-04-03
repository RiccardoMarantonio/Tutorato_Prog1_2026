// Package helpers fornisce utility comuni per testare esercizi Go
// che leggono da stdin e scrivono su stdout.
package helpers

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// RunWithInput esegue il package dell'esercizio con l'input fornito
// e restituisce l'output catturato.
//
// Usage:
//
//	out := helpers.RunWithInput(t, "./01_numeri_pari", "10\n")
//	if !strings.Contains(out, "0\n2\n4\n") {
//	    t.Errorf("output non atteso: %s", out)
//	}
func RunWithInput(t *testing.T, relPath string, input string) string {
	t.Helper()

	// Trova il percorso assoluto relativo al modulo
	dir := filepath.Join(".", relPath)

	cmd := exec.Command("go", "run", dir)
	cmd.Stdin = strings.NewReader(input)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("esecuzione fallita: %v\nstderr: %s", err, stderr.String())
	}

	return stdout.String()
}

// CaptureStdout cattura l'output di una funzione che scrive su os.Stdout.
// Restituisce una funzione da chiamare per ottenere la stringa catturata.
//
// Usage:
//
//	capture := helpers.CaptureStdout()
//	fmt.Println("ciao")
//	out := capture()
//	// out == "ciao\n"
func CaptureStdout() func() string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	return func() string {
		w.Close()
		os.Stdout = old

		var buf bytes.Buffer
		io.Copy(&buf, r)
		r.Close()
		return buf.String()
	}
}

// AssertEqual confronta due valori e fallisce il test se diversi.
func AssertEqual(t *testing.T, got, want interface{}, msg string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %v, want %v", msg, got, want)
	}
}

// AssertTrue fallisce il test se la condizione è falsa.
func AssertTrue(t *testing.T, condition bool, msg string) {
	t.Helper()
	if !condition {
		t.Errorf("%s: expected true, got false", msg)
	}
}

// AssertFalse fallisce il test se la condizione è vera.
func AssertFalse(t *testing.T, condition bool, msg string) {
	t.Helper()
	if condition {
		t.Errorf("%s: expected false, got true", msg)
	}
}

// AssertNoError fallisce il test se err non è nil.
func AssertNoError(t *testing.T, err error, msg string) {
	t.Helper()
	if err != nil {
		t.Errorf("%s: unexpected error: %v", msg, err)
	}
}

// AssertError fallisce il test se err è nil.
func AssertError(t *testing.T, err error, msg string) {
	t.Helper()
	if err == nil {
		t.Errorf("%s: expected error, got nil", msg)
	}
}

// AssertContains verifica che la stringa contenga il sottostringa.
func AssertContains(t *testing.T, s, substr string, msg string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("%s: expected %q to contain %q", msg, s, substr)
	}
}

// AssertSliceEqual confronta due slice di int.
func AssertSliceEqual(t *testing.T, got, want []int, msg string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: length mismatch: got %d, want %d", msg, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s: index %d: got %d, want %d", msg, i, got[i], want[i])
			return
		}
	}
}

// AssertSliceStringEqual confronta due slice di string.
func AssertSliceStringEqual(t *testing.T, got, want []string, msg string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: length mismatch: got %d, want %d (got=%v, want=%v)", msg, len(got), len(want), got, want)
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s: index %d: got %q, want %q", msg, i, got[i], want[i])
			return
		}
	}
}

// AssertFloatEqual confronta due float64 con tolleranza.
func AssertFloatEqual(t *testing.T, got, want float64, msg string) {
	t.Helper()
	tolerance := 0.01
	diff := got - want
	if diff < 0 {
		diff = -diff
	}
	if diff > tolerance {
		t.Errorf("%s: got %.2f, want %.2f", msg, got, want)
	}
}

// AssertMapEqual confronta due map[string]int.
func AssertMapEqual(t *testing.T, got, want map[string]int, msg string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: size mismatch: got %d entries, want %d", msg, len(got), len(want))
		return
	}
	for k, v := range want {
		if gv, ok := got[k]; !ok || gv != v {
			t.Errorf("%s: key %q: got %d, want %d", msg, k, gv, v)
			return
		}
	}
}

// NormalizeOutput rimuove spazi extra e newline finali per confronti robusti.
func NormalizeOutput(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return s
}
