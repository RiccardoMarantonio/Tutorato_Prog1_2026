# Numeri Pari

Scrivi un programma che legge un numero intero positivo `n` da standard input e stampa, uno per riga, tutti i numeri pari da `0` a `n` (incluso).

Il programma deve implementare una funzione:

```go
func IsEven(n int) bool
```

che restituisce `true` se `n` è pari, `false` altrimenti.

## Vincoli
- `n >= 0`
- Usa un ciclo `for` per iterare
- Chiama `IsEven` per ogni numero

## Esempio

Input:
```
10
```

Output:
```
0
2
4
6
8
10
```

## Suggerimento
- Usa `fmt.Scan` per leggere da stdin
- L'operatore `%` (modulo) restituisce il resto della divisione
