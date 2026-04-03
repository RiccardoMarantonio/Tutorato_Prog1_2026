# Somma Ricorsiva

Scrivi una funzione ricorsiva che calcola la somma dei numeri da 1 a n.

Il programma deve implementare:

```go
func SommaRicorsiva(n int) int
```

Se `n <= 0`, restituisce 0. Altrimenti restituisce `n + SommaRicorsiva(n-1)`.

## Vincoli
- La funzione deve essere ricorsiva (nessun ciclo `for`)
- `n >= 0`

## Esempio

Input:
```
5
```

Output:
```
Somma da 1 a 5: 15
```

## Suggerimento
- Caso base: `SommaRicorsiva(0) = 0`
- Passo ricorsivo: `SommaRicorsiva(n) = n + SommaRicorsiva(n-1)`
