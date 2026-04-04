# Contatore di Parole

Scrivi un programma che legge un testo da stdin e conta quante volte ciascuna parola appare. Usa una `map[string]int`.

Il testo puo contenere piu parole sulla stessa riga, separate da spazi. L'input termina con una riga vuota.

Il programma deve stampare tutte le parole con il loro conteggio alla fine dell'input.

## Esempio

Input:
```
il gatto e sul tetto il tetto e alto

```

Output:
```
il: 2
gatto: 1
e: 2
sul: 1
tetto: 2
alto: 1
```

## Vincoli
- Le parole sono separate da spazi
- Non usare mappe per ora... aspetta, questo esercizio e sulle mappe! Usa `map[string]int`
- L'input termina con una riga vuota

## Suggerimento
- Usa `bufio.Scanner` per leggere riga per riga
- Per ogni riga, usa `strings.Fields` per separare le parole
- In Go, le mappe si creano con `make(map[string]int)`
