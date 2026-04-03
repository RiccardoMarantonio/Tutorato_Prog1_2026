# Inversione e Palindromi

Scrivi un programma che lavora con slice e stringhe per verificare palindromi e invertire sequenze.

Il programma deve implementare le seguenti funzioni:

```go
func InvertiSlice(slice []int) []int
```
Restituisce una nuova slice con gli elementi in ordine inverso (la slice originale non deve essere modificata).

```go
func EPalindromo(s string) bool
```
Restituisce `true` se la stringa `s` è un palindromo (si legge uguale da sinistra e da destra). Il confronto deve essere case-insensitive e deve ignorare spazi e punteggiatura.

```go
func EPalindromoNumerico(numeri []int) bool
```
Restituisce `true` se la slice di numeri letta da sinistra è uguale a quella letta da destra.

## Vincoli
- Non usare librerie esterne oltre a `fmt`
- Per il palindromo di stringhe, converti tutto in minuscolo con una funzione helper
- La funzione `InvertiSlice` non deve modificare la slice originale

## Esempio

Input:
```
Anna
1 2 3 2 1
```

Output:
```
"Anna" è un palindromo: true
[1 2 3 2 1] è un palindromo numerico: true
Slice invertita: [1 2 3 2 1]
```

## Suggerimento
- In Go le stringhe sono sequenze di byte. Per questo esercizio puoi usare l'indicizzazione diretta se lavori solo con ASCII
- Per convertire un byte in minuscolo: se è tra 'A' e 'Z', aggiungi 32
