# Calcolatrice da Terminale

Scrivi un programma che legge da stdin un'espressione nella forma `a op b` (dove `a` e `b` sono interi e `op` e uno tra `+`, `-`, `*`, `/`) e stampa il risultato.

## Formato Input
Tre valori separati da spazio: `<numero> <operatore> <numero>`

## Regole
- Se l'operatore non e valido, stampa `"Operatore non valido"`
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
- Scrivi tutto il programma da zero (main e logica)
- Usa `switch` per gestire gli operatori
