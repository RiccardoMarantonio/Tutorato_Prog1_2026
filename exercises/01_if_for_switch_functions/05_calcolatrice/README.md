# Calcolatrice da Terminale

Scrivi un programma che legge da stdin un'espressione nella forma `a op b` (dove `a` e `b` sono interi e `op` e uno tra `+`, `-`, `*`, `/`) e stampa il risultato.

Il programma deve implementare la seguente funzione:

```go
func eseguiOp(a, b int, op rune) int
```
Esegue l'operazione tra `a` e `b` secondo l'operatore `op`. Usa `switch` per gestire i casi `+`, `-`, `*`, `/`. Per un operatore non valido, restituisce 0.

## Formato Input
Tre valori separati da spazio: `<numero> <operatore> <numero>`

## Regole
- Se l'operatore non e uno tra `+`, `-`, `*`, `/`, stampa `"Operatore non valido"`
- Se la divisione e per zero, stampa `"Divisione per zero"`
- Il risultato della divisione intera deve essere troncato verso zero (comportamento standard di Go)

## Esempi

Input:
```
5 + 3
```
Output:
```
8
```

Input:
```
10 / 0
```
Output:
```
Divisione per zero
```

Input:
```
7 % 3
```
Output:
```
Operatore non valido
```

## Vincoli
- Solo `fmt` come libreria
- La funzione `eseguiOp` deve usare `switch`
- Il main deve validare l'operatore prima di chiamare `eseguiOp`

## Suggerimento
- Leggi l'operatore come stringa con `fmt.Scan(&opStr)` e converti a rune con `[]rune(opStr)[0]`
- Nel main, valida l'operatore con un `if` prima di chiamare `eseguiOp`
