# Rotazione di una Slice

Scrivi un programma che legge una slice di numeri interi e un numero `k`, poi ruota la slice di `k` posizioni verso destra. Gli elementi che "escono" dalla fine rientrano dall'inizio.

Il programma deve implementare:

```go
func Ruota(slice []int, k int) []int
```
Restituisce una NUOVA slice ruotata. La slice originale NON deve essere modificata.

## Esempio

Input:
```
5
1 2 3 4 5
2
```

Output:
```
Originale: [1 2 3 4 5]
Ruotata:   [4 5 1 2 3]
```

## Vincoli
- Il primo numero e la lunghezza della slice
- `k` puo essere maggiore della lunghezza della slice
- La slice originale non deve essere modificata

## Suggerimento
- Usa `k % len(slice)` per gestire rotazioni maggiori della lunghezza
