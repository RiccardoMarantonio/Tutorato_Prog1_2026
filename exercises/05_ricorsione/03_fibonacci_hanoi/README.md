# Fibonacci con Memoizzazione e Torre di Hanoi

Questo esercizio copre due classici problemi ricorsivi.

## Parte 1: Fibonacci con Memoizzazione

Implementa una versione ottimizzata di Fibonacci che usa la memoizzazione per evitare calcoli ripetuti.

```go
func FibonacciMemo(n int) int
```
Calcola F(n) usando memoizzazione. Deve essere efficiente anche per n fino a 50.

```go
func ContaChiamate(n int) int
```
Quante chiamate ricorsive servono per calcolare F(n) SENZA memoizzazione? Restituisce questo numero.

## Parte 2: Torre di Hanoi

Implementa il risolutore della Torre di Hanoi. Dati `n` dischi e 3 pioli (A, B, C), stampa la sequenza di mosse per spostare tutti i dischi da A a C usando B come appoggio.

Regole:
- Si può spostare un solo disco alla volta
- Un disco più grande non può mai stare sopra un disco più piccolo

```go
func Hanoi(n int, sorgente, ausiliario, destinazione string) []string
```
Restituisce una slice di stringhe, ognuna descrive una mossa (es. `"Sposta disco da A a C"`).

## Vincoli
- FibonacciMemo deve usare una mappa o slice come cache
- Hanoi deve essere puramente ricorsivo
- `n` per Hanoi: 1 <= n <= 10

## Esempio

Input:
```
FIB 10
HANOI 3
```

Output:
```
Fibonacci(10) = 55
Chiamate senza memo: 177
Chiamate con memo: 19

Torre di Hanoi (3 dischi):
1. Sposta disco da A a C
2. Sposta disco da A a B
3. Sposta disco da C a B
4. Sposta disco da A a C
5. Sposta disco da B a A
6. Sposta disco da B a C
7. Sposta disco da A a C
Totale mosse: 7
```

## Suggerimento
- Per Hanoi: per spostare n dischi da A a C, sposta n-1 da A a B, poi il disco n da A a C, poi n-1 da B a C
- La memoizzazione si implementa con una `map[int]int` o `[]int` passata come parametro helper
