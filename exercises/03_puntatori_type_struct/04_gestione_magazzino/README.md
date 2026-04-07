# Gestione Magazzino Prodotti

Implementa un sistema di gestione magazzino con prodotti e movimenti.

Definisci le seguenti strutture:

```go
type Prodotto struct {
	Codice   string
	Nome     string
	Prezzo   float64
	Quantita int
}

type Magazzino struct {
	Prodotti map[string]Prodotto
}
```

Implementa le seguenti funzioni:

```go
func AggiungiProdotto(m *Magazzino, p Prodotto) error
```
Aggiunge un prodotto. Restituisce errore se il codice esiste già.

```go
func RegistraCarico(m *Magazzino, codice string, qta int) error
```
Aumenta la quantità di un prodotto. Errore se il prodotto non esiste o qta <= 0.

```go
func RegistraScarico(m *Magazzino, codice string, qta int) error
```
Diminuisce la quantità di un prodotto. Errore se il prodotto non esiste, qta <= 0, o quantità insufficiente.

```go
func ValoreMagazzino(m Magazzino) float64
```
Restituisce il valore totale del magazzino (somma di prezzo × quantità per ogni prodotto).

```go
func ProdottiSottoScorta(m Magazzino, soglia int) int
```
Restituisce il numero di prodotti con quantità inferiore alla soglia.

```go
func CercaProdotto(m Magazzino, codice string) (Prodotto, error)
```
Cerca un prodotto per codice. Errore se non trovato.

## Vincoli
- Tutti i movimenti devono essere validati prima di essere registrati
- I codici prodotto sono stringhe univoche
- Usa puntatori dove è necessario modificare le strutture

## Esempio

Input:
```
AGGIUNGI P001 Laptop 899.99 10
AGGIUNGI P002 Mouse 29.99 50
SCARICO P001 3
CARICO P002 20
VALORE
SOTTO_SCORTA 15
FINE
```

Output:
```
Prodotto P001 aggiunto
Prodotto P002 aggiunto
Scarico P001: 3 pezzi
Carico P002: 20 pezzi

Valore totale magazzino: 13199.20 EUR

Prodotti sotto scorta (soglia: 15): 1
```

## Suggerimento
- Puoi usare un ciclo per cercare prodotti per codice nella slice
