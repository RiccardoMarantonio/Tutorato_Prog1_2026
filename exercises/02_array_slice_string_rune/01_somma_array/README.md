# Somma o Prodotto di un Array

Scrivi un programma che legge 5 numeri interi da stdin, li memorizza in un array e poi chiede all'utente quale operazione applicare: somma o prodotto.

Il programma deve implementare le seguenti funzioni:

```go
func SommaArray(arr [5]int) int
```
Restituisce la somma di tutti gli elementi.

```go
func ProdottoArray(arr [5]int) int
```
Restituisce il prodotto di tutti gli elementi.

## Formato Input
Prima 5 numeri interi (gli elementi dell'array), poi una stringa `"somma"` o `"prodotto"` per scegliere l'operazione.

## Esempio 1

Input:
```
1 2 3 4 5
somma
```

Output:
```
Risultato: 15
```

## Esempio 2

Input:
```
1 2 3 4 5
prodotto
```

Output:
```
Risultato: 120
```

## Vincoli
- L'array ha dimensione fissa 5
- Usa un ciclo `for` per entrambe le operazioni
- Se l'operazione non è valida, stampa `"Operazione non valida"`
- Scrivi tutto il programma da zero
