# Somma di un Array

Scrivi un programma che legge 5 numeri interi da stdin, li memorizza in un array e ne calcola la somma.

Il programma deve implementare una funzione:

```go
func SommaArray(arr [5]int) int
```

che restituisce la somma di tutti gli elementi dell'array.

## Vincoli
- L'array ha dimensione fissa 5
- Usa un ciclo `for` per sommare gli elementi

## Esempio

Input:
```
1 2 3 4 5
```

Output:
```
Somma: 15
```

## Suggerimento
- Usa `fmt.Scan` con l'operatore `&` per leggere i valori nell'array
- Un array in Go ha dimensione fissa: `[5]int`
