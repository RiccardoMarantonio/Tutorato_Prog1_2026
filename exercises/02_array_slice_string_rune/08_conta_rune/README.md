# Conteggio Rune

Scrivi un programma che legge una stringa da stdin e conta:
- Il numero totale di rune (caratteri Unicode)
- Il numero di lettere (a-z, A-Z)
- Il numero di cifre (0-9)
- Il numero di spazi
- Il numero di altri caratteri (punteggiatura, simboli, emoji, ecc.)

Il programma deve implementare:

```go
func ContaRune(s string) map[string]int
```
Restituisce una mappa con le chiavi: `"lettere"`, `"cifre"`, `"spazi"`, `"altri"`.

## Esempio

Input:
```
Ciao 123!
```

Output:
```
Totale rune: 9
Lettere: 4
Cifre: 3
Spazi: 1
Altri: 1
```

## Vincoli
- La stringa puo contenere caratteri Unicode (accenti, emoji, ecc.)
- Usa `range` sulla stringa per iterare sulle rune, NON sui byte
- Scrivi tutto il programma da zero
