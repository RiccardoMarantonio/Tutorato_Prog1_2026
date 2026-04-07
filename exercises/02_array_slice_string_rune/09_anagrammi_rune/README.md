# Anagrammi con Rune

Date due stringhe da stdin, verifica se sono anagrammi (contengono gli stessi caratteri con le stesse frequenze).

Il programma deve implementare:

```go
func SonoAnagrammi(s1, s2 string) bool
```
Restituisce `true` se le due stringhe sono anagrammi. Il confronto deve essere case-insensitive.

```go
func FirmaRune(s string) []rune
```
Restituisce una slice di rune ordinata (firma) della stringa. Due stringhe sono anagrammi se e solo se le loro firme sono identiche.

## Esempi

Input:
```
listen silent
```
Output:
```
"listen" e "silent" sono anagrammi: true
```

Input:
```
hello world
```
Output:
```
"hello" e "world" sono anagrammi: false
```

## Vincoli
- Confronto case-insensitive
- Le stringhe possono contenere caratteri Unicode
- Per ordinare le rune, implementa un semplice bubble sort
- Scrivi tutto il programma da zero

## Suggerimento
- Ordina le rune e confronta le due slice
