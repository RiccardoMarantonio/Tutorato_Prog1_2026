# Anagrammi

Scrivi un programma che, dato un elenco di parole da stdin, raggruppa le parole che sono anagrammi tra loro (parole formate dagli stessi caratteri con le stesse frequenze).

Il programma deve implementare le seguenti funzioni:

```go
func FirmaAnagramma(parola string) string
```
Restituisce una "firma" della parola: i caratteri ordinati in ordine alfabetico. Due parole sono anagrammi se e solo se hanno la stessa firma.

```go
func RaggruppaAnagrammi(parole []string) [][]string
```
Raggruppa le parole per anagrammi usando una mappa. Ogni gruppo contiene parole che sono anagrammi tra loro.

```go
func ContaAnagrammi(parole []string) map[string]int
```
Restituisce una mappa che associa ogni firma al numero di parole che la condividono.

## Vincoli
- Il primo numero letto da stdin indica quante parole seguiranno
- Confronto case-insensitive (converti tutto in minuscolo)
- Solo lettere dell'alfabeto inglese (a-z)
- Per ordinare i caratteri di una stringa, convertila in una slice di byte/rune e ordinala con un ciclo

## Esempio

Input:
```
6
listen
silent
enlist
hello
world
drowl
```

Output:
```
Gruppi di anagrammi:
[listen silent enlist]
[hello]
[world drowl]

Firme più comuni:
eilnst: 3
dlorw: 2
ehllo: 1
```

## Suggerimento
- Puoi ordinare i caratteri di una stringa convertendola in `[]byte` e usando un semplice bubble sort
- Usa `map[string][]string` per raggruppare: la chiave è la firma, il valore è la lista di parole
