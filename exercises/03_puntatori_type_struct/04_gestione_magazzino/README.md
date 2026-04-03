# Gestione Magazzino Prodotti

Implementa un sistema di gestione magazzino con prodotti, categorie e movimenti.

Definisci le seguenti strutture:

```go
type Prodotto struct {
	Codice    string
	Nome      string
	Categoria string
	Prezzo    float64
	Quantita  int
}

type Movimento struct {
	CodiceProdotto string
	Tipo           string // "CARICO" o "SCARICO"
	Quantita       int
}

type Magazzino struct {
	Prodotti  []Prodotto
	Movimenti []Movimento
}
```

Implementa le seguenti funzioni:

```go
func AggiungiProdotto(m *Magazzino, p Prodotto) error
```
Aggiunge un prodotto. Restituisce errore se il codice esiste già.

```go
func RegistraMovimento(m *Magazzino, mov Movimento) error
```
Registra un movimento e aggiorna la quantità del prodotto. Restituisce errore se il prodotto non esiste, lo scarico richiede più pezzi di quelli disponibili, o la quantità è <= 0.

```go
func ValoreMagazzino(m Magazzino) float64
```
Restituisce il valore totale del magazzino (somma di prezzo × quantità per ogni prodotto).

```go
func ProdottiSottoScorta(m Magazzino, soglia int) []Prodotto
```
Restituisce i prodotti con quantità inferiore alla soglia.

```go
func RiepilogoCategoria(m Magazzino) map[string]int
```
Restituisce una mappa che associa ogni categoria al numero totale di pezzi in magazzino.

## Vincoli
- Tutti i movimenti devono essere validati prima di essere registrati
- I codici prodotto sono stringhe univoche
- Usa puntatori dove è necessario modificare le strutture

## Esempio

Input:
```
AGGIUNGI P001 Laptop Elettronica 899.99 10
AGGIUNGI P002 Mouse Elettronica 29.99 50
SCARICO P001 3
CARICO P002 20
VALORE
SOTTO_SCORTA 10
RIEPILOGO
FINE
```

Output:
```
Prodotto P001 aggiunto
Prodotto P002 aggiunto
Scarico P001: 3 pezzi
Carico P002: 20 pezzi

Valore totale magazzino: 13199.20 EUR

Prodotti sotto scorta (soglia: 10):
Nessuno

Riepilogo per categoria:
Elettronica: 77
```

## Suggerimento
- Puoi usare un ciclo per cercare prodotti per codice nella slice
- Gestisci gli errori restituendoli con `fmt.Errorf("messaggio")`
