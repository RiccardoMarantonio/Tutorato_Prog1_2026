# Fibonacci con Memoizzazione

Implementa una versione ottimizzata di Fibonacci che usa la memoizzazione per evitare calcoli ripetuti.

La sequenza di Fibonacci e definita come:
- F(0) = 0
- F(1) = 1
- F(n) = F(n-1) + F(n-2) per n > 1

Il programma deve implementare le seguenti funzioni:

```go
func FibonacciMemo(n int) int
```
Calcola F(n) usando memoizzazione. Deve essere efficiente anche per n fino a 50.

```go
func ContaChiamate(n int) int
```
Quante chiamate ricorsive servono per calcolare F(n) SENZA memoizzazione? Restituisce questo numero.

## Vincoli
- `FibonacciMemo` deve usare una mappa o slice come cache
- `ContaChiamate` deve essere puramente ricorsivo (senza memoizzazione)
- `n` puo arrivare fino a 50 per `FibonacciMemo` e fino a 40 per `ContaChiamate`

## Esempio

Input:
```
10
```

Output:
```
Fibonacci(10) = 55
Chiamate senza memo: 177
Chiamate con memo: 19
```

## Suggerimento
- La memoizzazione si implementa con una `map[int]int` o `[]int` passata come parametro helper
- Senza memoizzazione, F(40) richiede oltre 300 milioni di chiamate
- Con memoizzazione, F(50) richiede solo ~99 chiamate
