# Voti degli Studenti

Scrivi un programma che gestisce i voti di una classe di studenti usando le mappe.

Il programma deve implementare le seguenti funzioni:

```go
type Studente struct {
	Nome string
	Voti []int
}

func AggiungiVoto(registro map[string]*Studente, nome string, voto int) error
```
Aggiunge un voto allo studente. Se lo studente non esiste, lo crea. Restituisce errore se il voto non è tra 1 e 30.

```go
func MediaStudente(registro map[string]*Studente, nome string) (float64, error)
```
Calcola la media dei voti di uno studente. Restituisce errore se lo studente non esiste o non ha voti.

```go
func MigliorPeggiore(registro map[string]*Studente) (string, string, error)
```
Restituisce i nomi dello studente con la media più alta e di quello con la media più bassa. Restituisce errore se il registro è vuoto.

```go
func DistribuzioneVoti(registro map[string]*Studente) map[string]int
```
Restituisce una mappa con la distribuzione dei voti per fascia:
- `"Insufficiente"` (1-17)
- `"Sufficiente"` (18-23)
- `"Buono"` (24-27)
- `"Ottimo"` (28-30)

## Vincoli
- Usa `map[string]*Studente` come registro (nome → puntatore a studente)
- I voti sono interi da 1 a 30
- Le medie devono essere stampate con 2 decimali

## Esempio

Input:
```
AGGIUNGI Mario 25
AGGIUNGI Mario 28
AGGIUNGI Luigi 18
AGGIUNGI Anna 30
AGGIUNGI Anna 30
MEDIA Mario
MIGLIOR_PEGGIORE
DISTRIBUZIONE
FINE
```

Output:
```
Voto aggiunto per Mario: 25
Voto aggiunto per Mario: 28
Voto aggiunto per Luigi: 18
Voto aggiunto per Anna: 30
Voto aggiunto per Anna: 30

Media di Mario: 26.50

Migliore: Anna (30.00)
Peggiore: Luigi (18.00)

Distribuzione voti:
Insufficiente: 0
Sufficiente: 1
Buono: 2
Ottimo: 2
```

## Suggerimento
- Per la distribuzione, scorri tutti gli studenti e tutti i loro voti
