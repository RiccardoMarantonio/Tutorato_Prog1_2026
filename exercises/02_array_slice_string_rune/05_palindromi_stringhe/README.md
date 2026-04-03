# Palindromi di Stringhe

Scrivi un programma che verifica se una stringa letta da stdin e un palindromo (si legge uguale da sinistra e da destra).

Il confronto deve essere **case-insensitive** e deve **ignorare gli spazi**.

Il programma deve implementare la funzione:

```go
func EPalindromo(s string) bool
```

## Esempi

Input:
```
Anna
```
Output:
```
"Anna" e un palindromo: true
```

Input:
```
hello
```
Output:
```
"hello" e un palindromo: false
```

## Vincoli
- Solo caratteri ASCII (a-z, A-Z, spazio)
- Per convertire in minuscolo: se un byte e tra 'A' e 'Z', aggiungi 32
- Scrivi tutto il programma da zero

## Suggerimento
- Prima pulisci la stringa: rimuovi spazi e converti in minuscolo
- Poi confronta i caratteri da entrambi i lati verso il centro
