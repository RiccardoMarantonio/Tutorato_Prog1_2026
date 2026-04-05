# Wordle Semplificato

Implementa una versione semplificata di Wordle. La parola segreta e fissata a `"crane"` (5 lettere).

Il giocatore ha **6 tentativi**. Dopo ogni tentativo:
- Se la lettera e nella posizione corretta → `"VERDE"`
- Se la lettera esiste nella parola ma in posizione sbagliata → `"GIALLO"`
- Se la lettera non esiste nella parola → `"GRIGIO"`

Per ogni tentativo, stampa il feedback lettera per lettera.

Il programma deve implementare le seguenti funzioni:

```go
func Feedback(tentativo, segreto string) []rune
```
Restituisce una slice di rune dove ogni elemento e:
- `'V'` (verde) se la lettera e corretta nella posizione corretta
- `'G'` (giallo) se la lettera esiste ma posizione sbagliata
- `'X'` (grigio) se la lettera non esiste

```go
func StampaFeedback(feedback []rune)
```
Stampa il feedback in formato leggibile: `V=VERDE, G=GIALLO, X=GRIGIO`.

## Esempio

Input:
```
crane
```

Output:
```
Tentativo 1/6: crane
V V V V V
Hai indovinato! Vittoria in 1 tentativi!
```

Input:
```
trace
crane
```

Output:
```
Tentativo 1/6: trace
V X G V V
Tentativo 2/6: crane
V V V V V
Hai indovinato! Vittoria in 2 tentativi!
```

## Vincoli
- La parola segreta e `"crane"` (5 lettere)
- Il giocatore ha esattamente 6 tentativi
- I tentativi devono essere esattamente 5 lettere
- Non usare mappe
- Scrivi tutto il programma da zero

## Suggerimento
- Per le lettere gialle: conta le occorrenze nel segreto e confronta con quelle gia assegnate come verdi
