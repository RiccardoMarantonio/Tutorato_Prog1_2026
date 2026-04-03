# Fibonacci Ricorsivo

Scrivi un programma che calcola l'n-esimo numero di Fibonacci usando la ricorsione.

La sequenza di Fibonacci è definita come:
- F(0) = 0
- F(1) = 1
- F(n) = F(n-1) + F(n-2) per n > 1

Il programma deve implementare le seguenti funzioni:

```go
func Fibonacci(n int) int
```
Restituisce l'n-esimo numero di Fibonacci (versione ricorsiva semplice).

```go
func FibonacciFinoA(n int) []int
```
Restituisce una slice con tutti i numeri di Fibonacci da F(0) a F(n).

## Vincoli
- `0 <= n <= 30` (la versione ricorsiva semplice è lenta per n grandi)
- La funzione `Fibonacci` deve essere puramente ricorsiva
- `FibonacciFinoA` può usare un ciclo per chiamare `Fibonacci`

## Esempio

Input:
```
10
```

Output:
```
Fibonacci(10) = 55
Sequenza: [0 1 1 2 3 5 8 13 21 34 55]
```

## Suggerimento
- La ricorsione semplice per Fibonacci ha complessità O(2^n)
- Per n=30 il calcolo potrebbe richiedere qualche secondo
- Non usare memoizzazione in questo esercizio
