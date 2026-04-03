# Gestione Biblioteca

Implementa un sistema completo per la gestione di una biblioteca. Questo esercizio ripassa tutti i concetti visti: if/for/switch, funzioni, array/slice, stringhe, puntatori, struct e mappe.

## Strutture Dati

```go
type Libro struct {
	ISBN        string
	Titolo      string
	Autore      string
	Anno        int
	Genere      string
	Disponibile bool
}

type Prestito struct {
	ISBN           string
	Utente         string
	DataPrestito   string // "GG/MM/AAAA"
	DataRestituzione string // "", vuota se non ancora restituito
}

type Biblioteca struct {
	Libri    map[string]*Libro // ISBN → Libro
	Prestiti []Prestito
}
```

## Funzioni Richieste

```go
func NuovaBiblioteca() *Biblioteca
```
Crea e restituisce una nuova biblioteca vuota.

```go
func AggiungiLibro(b *Biblioteca, l Libro) error
```
Aggiunge un libro alla biblioteca. Errore se l'ISBN esiste già.

```go
func CercaLibro(b Biblioteca, termine string) []Libro
```
Cerca libri per titolo o autore (case-insensitive, contiene il termine).

```go
func RegistraPrestito(b *Biblioteca, isbn, utente, data string) error
```
Registra un prestito. Errore se il libro non esiste o non è disponibile.

```go
func RegistraRestituzione(b *Biblioteca, isbn, data string) error
```
Registra la restituzione di un libro. Errore se il libro non esiste o non è in prestito.

```go
func LibriPerGenere(b Biblioteca) map[string]int
```
Restituisce una mappa che associa ogni genere al numero di libri.

```go
func PrestitiAperti(b Biblioteca) []Prestito
```
Restituisce tutti i prestiti non ancora chiusi (dataRestituzione vuota).

```go
func Statistiche(b Biblioteca) map[string]int
```
Restituisce statistiche: totale libri, disponibili, in prestito, generi diversi.

## Comandi da stdin

Il programma legge comandi da stdin fino a `FINE`:

```
AGGIUNGI <isbn> <titolo> <autore> <anno> <genere>
CERCA <termine>
PRESTITO <isbn> <utente> <data>
RESTITUZIONE <isbn> <data>
CATALOGO
PER_GENERI
PRESTITI_APERTI
STATISTICHE
FINE
```

## Vincoli
- Solo librerie standard Go: `fmt`
- Usa puntatori dove appropriato (modifica struct nella mappa)
- Validazione input di base
- Ricerca case-insensitive implementata manualmente

## Esempio

Input:
```
AGGIUNGI 978-001 Il Nome Della Rosa Umberto Eco 1980 Narrativa
AGGIUNGI 978-002 1984 George Orwell 1949 Narrativa
AGGIUNGI 978-003 Clean Code Robert Martin 2008 Informatica
PRESTITO 978-001 Mario 15/01/2025
PER_GENERI
PRESTITI_APERTI
STATISTICHE
FINE
```

Output:
```
Libro aggiunto: Il Nome Della Rosa
Libro aggiunto: 1984
Libro aggiunto: Clean Code
Prestito registrato: 978-001 a Mario

Libri per genere:
Narrativa: 2
Informatica: 1

Prestiti aperti:
978-001 "Il Nome Della Rosa" → Mario (dal 15/01/2025)

Statistiche:
Totale libri: 3
Disponibili: 2
In prestito: 1
Generi diversi: 2
```

## Suggerimenti
- `map[string]*Libro` permette di modificare il campo `Disponibile` senza reinserire il libro
- Per la ricerca case-insensitive, converti tutto in minuscolo confrontando i byte
- I prestiti aperti sono quelli con `DataRestituzione == ""`
