# Contatore di Parole

Scrivi un programma che legge parole da stdin (una per riga) e conta quante volte ciascuna parola appare. Usa una `map[string]int`.

Il programma deve stampare tutte le parole con il loro conteggio alla fine dell'input (riga vuota per terminare).

## Esempio

Input:
```
ciao
mondo
ciao
go

```

Output:
```
ciao: 2
mondo: 1
go: 1
```

## Suggerimento
- In Go, le mappe si creano con `make(map[string]int)`
- Per verificare se una chiave esiste: `val, ok := m[key]`
