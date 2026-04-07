# Elezioni con Mappe

Scrivi un programma che gestisce il risultato di un'elezione con più candidati e sezioni elettorali.

Il programma deve implementare le seguenti funzioni:

```go
type Sezione struct {
	Nome     string
	Voti     map[string]int // candidato → voti
	Elettori int
	Votanti  int
}

type Elezione struct {
	Candidati []string
	Sezioni   map[string]*Sezione // nome sezione → sezione
}

func RegistraVoto(e *Elezione, sezione, candidato string) error
```
Registra un voto per un candidato in una sezione. Restituisce errore se la sezione non esiste, il candidato non è valido, o i votanti hanno raggiunto gli elettori.

```go
func RisultatiNazionali(e Elezione) map[string]int
```
Restituisce il totale dei voti per ogni candidato su tutte le sezioni.

```go
func Affluenza(e Elezione) map[string]float64
```
Restituisce una mappa con la percentuale di affluenza per ogni sezione (votanti / elettori * 100).

```go
func Vincitore(e Elezione) (string, error)
```
Restituisce il nome del candidato con più voti a livello nazionale. Errore se non ci sono voti.

```go
func Ballottaggio(e Elezione) (string, string, error)
```
Restituisce i nomi dei due candidati più votati (per un eventuale ballottaggio). Errore se ci sono meno di 2 candidati.

## Vincoli
- I candidati vengono registrati all'inizio
- Le sezioni vengono create con un numero di elettori
- Ogni voto incrementa il contatore dei votanti della sezione
- Non si può superare il numero di elettori

## Esempio

Input:
```
CANDIDATI Rossi Bianchi Verdi
SEZIONE Centro 1000
SEZIONE Nord 800
VOTO Centro Rossi
VOTO Centro Bianchi
VOTO Nord Verdi
RISULTATI
AFFLUENZA
VINCITORE
FINE
```

Output:
```
Candidati registrati: [Rossi Bianchi Verdi]
Sezione Centro creata (1000 elettori)
Sezione Nord creata (800 elettori)
Voto registrato: Centro → Rossi
Voto registrato: Centro → Bianchi
Voto registrato: Nord → Verdi

=== Risultati Nazionali ===
Rossi: 1
Bianchi: 1
Verdi: 1

=== Affluenza ===
Centro: 0.20%
Nord: 0.13%

Vincitore: Rossi (1 voti)
```

## Suggerimento
- Per il ballottaggio, puoi ordinare i candidati per voti usando una slice temporanea
- L'affluenza è `float64(votanti) / float64(elettori) * 100`
