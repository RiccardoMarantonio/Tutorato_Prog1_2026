# Massimo e Minimo di una Slice

Scrivi un programma che legge una sequenza di numeri interi da stdin e trova il valore massimo e minimo.

Il programma deve implementare le seguenti funzioni:

```go
func TrovaMax(slice []int) int
```
Restituisce il valore massimo nella slice. Se vuota, restituisci 0.

```go
func TrovaMin(slice []int) int
```
Restituisce il valore minimo nella slice. Se vuota, restituisci 0.

```go
func TrovaMaxMin(slice []int) (int, int)
```
Restituisce sia il massimo che il minimo in una sola funzione.

## Esempio

Input:
```
5
10 3 7 1 9
```

Output:
```
Max: 10
Min: 1
```

## Vincoli
- Il primo numero letto indica quanti elementi seguiranno
- La slice deve essere creata con `make([]int, n)`
