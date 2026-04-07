# Gestione Biblioteca

Implementa un sistema completo per la gestione di una biblioteca. Questo esercizio ripassa tutti i concetti visti: if/for/switch, funzioni, array/slice, stringhe, puntatori, struct e mappe.

## Strutture Dati

Definisci tre struct:
- **Libro**: con ISBN, titolo, autore, anno, genere e un flag per la disponibilita
- **Prestito**: con ISBN del libro, nome utente, data del prestito e data di restituzione (vuota se non ancora restituito)
- **Biblioteca**: contiene una mappa di libri (indicizzata per ISBN) e una slice di prestiti

## Funzionalita da Implementare

### Gestione Libri
- **Creare una nuova biblioteca** vuota
- **Aggiungere un libro** alla biblioteca. L'ISBN deve essere unico: se esiste gia, restituisci un errore
- **Cercare libri** per titolo o autore. La ricerca deve essere case-insensitive e trovare libri che contengono il termine cercato (non devono essere uguali)
- **Visualizzare il catalogo** completo

### Gestione Prestiti
- **Registrare un prestito**: associa un libro a un utente in una data specifica. Il libro deve esistere ed essere disponibile. Se il libro non esiste o non e disponibile, restituisci un errore
- **Registrare una restituzione**: imposta la data di restituzione per un prestito esistente. Il libro deve esistere ed essere in prestito. Se non esiste o non e in prestito, restituisci un errore
- **Visualizzare i prestiti aperti**: mostra tutti i prestiti ancora aperti (senza data di restituzione)

### Statistiche
- **Contare i libri per genere**: restituisce una mappa che associa ogni genere al numero di libri in quel genere
- **Mostrare statistiche generali**: totale libri, libri disponibili, libri in prestito, numero di generi diversi

## Formato Input

Il programma legge comandi da stdin fino a `FINE`. Ogni comando e su una riga separata:

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
- Usa solo le librerie standard Go necessarie
- Usa puntatori dove appropriato (per modificare i campi delle struct nella mappa)
- Implementa la ricerca case-insensitive manualmente (confrontando i byte in minuscolo)
- Gestisci gli errori in modo appropriato (ISBN duplicato, libro non trovato, libro non disponibile, ecc.)

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

Output atteso:
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
- Usa `map[string]*Libro` per poter modificare il campo `Disponibile` senza dover reinserire il libro nella mappa
- Per la ricerca case-insensitive, converti sia il testo cercato che i campi titolo/autore in minuscolo prima di confrontare
- I prestiti aperti sono quelli con `DataRestituzione` vuota (stringa vuota)
