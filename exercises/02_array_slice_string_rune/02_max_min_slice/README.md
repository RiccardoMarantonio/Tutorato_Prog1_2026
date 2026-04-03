# Massimo e Minimo di una Slice

Scrivi un programma che legge una sequenza di numeri interi da stdin e trova il valore massimo e minimo.

Il programma deve implementare le seguenti funzioni:

```go
func TrovaMax(slice []int) int
```
Restituisce il valore massimo nella slice.

```go
func TrovaMin(slice []int) int
```
Restituisce il valore minimo nella slice.

```go
func TrovaMaxMin(slice []int) (int, int)
```
Restituisce sia il massimo che il minimo in una sola funzione.

## Vincoli
- Il primo numero letto da stdin indica quanti elementi seguiranno
- La slice deve essere creata con `make([]int, n)`
- Se la slice è vuota, restituisci `0` per entrambi

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

## Suggerimento
- In Go le slice sono dinamiche, a differenza degli array
- Puoi usare `make([]int, n)` per creare una slice di dimensione `n`
- Inizializza max e min con il primo elemento della slice
