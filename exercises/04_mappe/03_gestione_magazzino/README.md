# Gestione Magazzino con Mappe

Implementa un sistema di gestione magazzino utilizzando le **mappe** per memorizzare i prodotti. La chiave della mappa sarà il codice del prodotto.

Definisci la struttura:

```go
type Prodotto struct {
	Nome     string
	Prezzo   float64
	Quantita int
}
```

Il magazzino sara rappresentato come:

```go
type Magazzino map[string]Prodotto
```

Implementa le seguenti funzioni:

```go
func AggiungiProdotto(m Magazzino, codice string, p Prodotto) error
```
Aggiunge un prodotto alla mappa. Restituisce errore se il codice esiste gia.

```go
func RegistraCarico(m Magazzino, codice string, qta int) error
```
Aumenta la quantita di un prodotto. Errore se il prodotto non esiste o qta <= 0.

```go
func RegistraScarico(m Magazzino, codice string, qta int) error
```
Diminuisce la quantita di un prodotto. Errore se il prodotto non esiste, qta <= 0, o quantita insufficiente.

```go
func ValoreMagazzino(m Magazzino) float64
```
Restituisce il valore totale del magazzino (somma di prezzo x quantita per ogni prodotto).

```go
func ProdottiSottoScorta(m Magazzino, soglia int) int
```
Restituisce il numero di prodotti con quantita inferiore alla soglia.

```go
func CercaProdotto(m Magazzino, codice string) (Prodotto, error)
```
Cerca un prodotto per codice nella mappa. Errore se non trovato.

```go
func NumeroProdotti(m Magazzino) int
```
Restituisce il numero totale di prodotti nel magazzino.

## Esempio

Input:
```
AGGIUNGI P001 Laptop 899.99 10
AGGIUNGI P002 Mouse 29.99 50
SCARICO P001 3
CARICO P002 20
VALORE
SOTTO_SCORTA 15
NUMERO
CERCA P001
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

Numero prodotti: 2

Trovato: P001 - Laptop - 899.99 EUR - Quantita: 7
```

## Vincoli
- Usa una `map[string]Prodotto` per il magazzino
- Tutti i movimenti devono essere validati prima di essere registrati
- I codici prodotto sono stringhe univoche (chiavi della mappa)

## Suggerimento
- Per verificare se un prodotto esiste nella mappa: `if prod, ok := m[codice]; ok { ... }`
- Per modificare il prodotto nella mappa, devi riassegnarlo: `m[codice] = prod`
