# Torre di Hanoi

Implementa il risolutore della Torre di Hanoi. Dati `n` dischi e 3 pioli (A, B, C), calcola la sequenza di mosse per spostare tutti i dischi da A a C usando B come appoggio.

Regole:
- Si puo spostare un solo disco alla volta
- Un disco piu grande non puo mai stare sopra un disco piu piccolo

Il programma deve implementare la seguente funzione:

```go
func Hanoi(n int, sorgente, ausiliario, destinazione string) []string
```
Restituisce una slice di stringhe, ognuna descrive una mossa (es. `"Sposta disco da A a C"`).

## Vincoli
- La funzione deve essere puramente ricorsiva
- `1 <= n <= 15`
- Il numero di mosse per n dischi e sempre `2^n - 1`

## Esempio

Input:
```
3
```

Output:
```
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
- Per spostare n dischi da A a C:
  1. Sposta n-1 dischi da A a B (usando C come appoggio)
  2. Sposta il disco n da A a C
  3. Sposta n-1 dischi da B a C (usando A come appoggio)
- Caso base: n=1, sposta direttamente da sorgente a destinazione
