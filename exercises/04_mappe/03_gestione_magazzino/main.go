package main

// Scrivi qui il tuo programma.
// Gestisci un magazzino prodotti usando una map[string]Prodotto.
// Comandi: AGGIUNGI <codice> <nome> <prezzo> <quantita>, CARICO <codice> <qta>,
// SCARICO <codice> <qta>, VALORE, SOTTO_SCORTA <soglia>, NUMERO, CERCA <codice>, FINE.
// Input: sequenza di comandi, uno per riga, terminata da FINE
// Output: messaggi di conferma per ogni operazione

type Prodotto struct {
	Nome     string
	Prezzo   float64
	Quantita int
}

type Magazzino map[string]Prodotto

func AggiungiProdotto(m Magazzino, codice string, p Prodotto) error {
	return nil
}

func RegistraCarico(m Magazzino, codice string, qta int) error {
	return nil
}

func RegistraScarico(m Magazzino, codice string, qta int) error {
	return nil
}

func ValoreMagazzino(m Magazzino) float64 {
	return 0
}

func ProdottiSottoScorta(m Magazzino, soglia int) int {
	return 0
}

func CercaProdotto(m Magazzino, codice string) (Prodotto, error) {
	return Prodotto{}, nil
}

func NumeroProdotti(m Magazzino) int {
	return 0
}

func main() {}
