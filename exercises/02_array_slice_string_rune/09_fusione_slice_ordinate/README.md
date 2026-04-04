# Fusione di Slice Ordinate

Scrivi un programma che lavora con slice ordinate per unirle e analizzarle.

Il programma deve implementare le seguenti funzioni:

```go
func MergeSorted(a, b []int) []int
```
Date due slice gia ordinate, restituisce una nuova slice ordinata con tutti gli elementi.

```go
func RimuoviDuplicati(slice []int) []int
```
Data una slice ordinata, restituisce una nuova slice senza duplicati.

```go
func Intersezione(a, b []int) []int
```
Date due slice ordinate, restituisce gli elementi comuni (senza duplicati).

## Esempio

Input:
```
5
1 3 5 7 9
6
2 3 5 8 9 10
```

Output:
```
Merge: [1 2 3 3 5 5 7 8 9 9 10]
Senza duplicati: [1 2 3 5 7 8 9 10]
Intersezione: [3 5 9]
```

## Vincoli
- Le slice in input sono gia ordinate
- Non usare `sort` o librerie esterne
- Non concatenare le due slice e poi ordinare: sfrutta il fatto che sono gia ordinate

## Suggerimento
- Per MergeSorted, usa due indici che avanzano simultaneamente nelle due slice
- Per Intersezione, avanza l'indice della slice con il valore minore
