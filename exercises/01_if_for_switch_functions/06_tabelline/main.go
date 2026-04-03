package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// RigaTabellina restituisce una singola riga della tabellina.
// Esempio: RigaTabellina(2, 3) → "2 x 3 = 6"
func RigaTabellina(base int, molt int) string {
	// TODO: implementare
	return ""
}

// StampaTabelline stampa le tabelline da 1 a n.
func StampaTabelline(n int) {
	// TODO: implementare
}

func main() {
	// TODO:
	// 1. Leggi un numero n da stdin
	// 2. Per ogni base da 1 a n:
	//    - Per ogni moltiplicatore da 1 a 10:
	//      - Stampa RigaTabellina(base, moltiplicatore)
	//    - Stampa una riga vuota dopo ogni tabellina
	_ = fmt.Scan // placeholder
}

// Helper per i test (non modificare)
func catturaOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	r.Close()
	return buf.String()
}

func contaRighe(s string) int {
	count := 0
	for _, c := range s {
		if c == '\n' {
			count++
		}
	}
	return count
}
