# Rotazione di una Stringa

Scrivi un programma che legge una stringa da stdin e un numero `k`, poi ruota la stringa di `k` posizioni verso destra. I caratteri che "escono" dalla fine rientrano dall'inizio.

La stringa puo contenere caratteri Unicode (accenti, emoji, ecc.), quindi **non puoi usare l'indicizzazione diretta sulla string**. Devi convertire la stringa in una slice di `rune`, ruotarla, e riconvertirla in stringa.

Il programma deve implementare le seguenti funzioni:

```go
func RuotaSlice(slice []int, k int) []int
```
Ruota una slice di interi di `k` posizioni verso destra (riutilizza la funzione dell'esercizio precedente sulla rotazione di slice).

```go
func RuotaStringa(s string, k int) string
```
Ruota una stringa di `k` posizioni verso destra. Deve convertire la stringa in `[]rune`, usare `RuotaSlice` per ruotare, e riconvertire in stringa.

## Esempio

Input:
```
ciao
2
```

Output:
```
Originale: ciao
Ruotata:   aoci
```

Input (con Unicode):
```
cafe\u0301
1
```

Output:
```
Originale: cafe\u0301
Ruotata:   \u0301cafe
```

## Vincoli
- La stringa puo contenere caratteri Unicode
- `k` puo essere maggiore della lunghezza della stringa
- **Devi riutilizzare la funzione `RuotaSlice` dell'esercizio precedente**
- Non usare l'indicizzazione diretta sulla string (`s[i]`) per accedere ai caratteri

## Suggerimento
- Per ruotare le rune, puoi riadattare `RuotaSlice` oppure scrivere una versione specifica per rune
