# Fusione di Slice Ordinate

Scrivi un programma che lavora con slice ordinate per unirle e analizzarle.

Il programma deve implementare le seguenti funzioni:

```go
func MergeSorted(a, b []int) []int
```
Date due slice già ordinate in modo crescente, restituisce una nuova slice ordinata che contiene tutti gli elementi di entrambe. La funzione deve essere efficiente: non basta concatenare e ordinare, ma deve sfruttare il fatto che le slice sono già ordinate (come nel merge del MergeSort).

```go
func RimuoviDuplicati(slice []int) []int
```
Data una slice ordinata, restituisce una nuova slice senza elementi duplicati, mantenendo l'ordine.

```go
func Intersezione(a, b []int) []int
```
Date due slice ordinate, restituisce una slice con gli elementi comuni a entrambe (senza duplicati).

```go
func ContaOccorrenze(slice []int, valore int) int
```
Conta quante volte `valore` appare nella slice.

## Vincoli
- Tutte le slice in input sono già ordinate
- Non usare `sort` o librerie esterne
- Le funzioni devono essere efficienti O(n) dove possibile

## Esempio

Input:
```
Slice A: 1 3 5 7 9
Slice B: 2 3 5 8 9 10
```

Output:
```
Merge: [1 2 3 3 5 5 7 8 9 9 10]
Merge senza duplicati: [1 2 3 5 7 8 9 10]
Intersezione: [3 5 9]
Occorrenze di 5: 2
```

## Suggerimento
- Per `MergeSorted`, usa due indici che avanzano nelle due slice simultaneamente
- Per `Intersezione`, se entrambe le slice sono ordinate puoi avanzare l'indice della slice con il valore minore
