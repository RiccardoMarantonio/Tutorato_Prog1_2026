# Espressioni Aritmetiche con Parentesi

Scrivi un programma che valuta espressioni aritmetiche contenute in stringhe. Le espressioni possono contenere numeri interi, le quattro operazioni (+, -, *, /) e parentesi tonde.

Il programma deve implementare le seguenti funzioni:

```go
func ValutaEspressione(expr string) (float64, error)
```
Valuta un'espressione aritmetica e restituisce il risultato. Supporta:
- Numeri interi positivi
- Operazioni: +, -, *, /
- Parentesi tonde `()`
- Spazi (da ignorare)

```go
func ContaOperatori(expr string) map[string]int
```
Conta quanti operatori di ogni tipo ci sono nell'espressione.

```go
func ProfonditaParentesi(expr string) int
```
Restituisce la massima profondita di annidamento delle parentesi.

```go
func EValida(expr string) bool
```
Verifica se l'espressione e sintatticamente valida (parentesi bilanciate, nessun operatore consecutivo, ecc.).

## Vincoli
- Puoi assumere che i numeri siano interi senza segno
- La divisione per zero deve restituire errore
- Non usare `eval` o librerie di parsing
- Implementa un parser ricorsivo semplice che gestisce la precedenza degli operatori

## Esempio

Input:
```
(3 + 5) * 2
10 / (2 + 3)
((1 + 2) * (3 + 4))
```

Output:
```
(3 + 5) * 2 = 16.00
10 / (2 + 3) = 2.00
((1 + 2) * (3 + 4)) = 21.00

Operatori:
+: 3
*: 2
/: 1

Massima profondita parentesi: 2
```

## Suggerimento
- Un approccio ricorsivo: `Espressione = Termine {+|- Termine}` e `Termine = Fattore {*|/ Fattore}` e `Fattore = Numero | (Espressione)`
- Puoi prima tokenizzare la stringa (separare numeri e operatori), poi valutare
- Per le parentesi bilanciate: conta le `(` aperte e `)` chiuse
