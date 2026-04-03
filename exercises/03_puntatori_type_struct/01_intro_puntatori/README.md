# Introduzione ai Puntatori

Scrivi un programma che dimostra il funzionamento base dei puntatori in Go.

Il programma deve implementare le seguenti funzioni:

```go
func Scambia(a, b *int)
```
Scambia i valori di due variabili intere usando i puntatori. Dopo la chiamata, le variabili originali devono essere modificate.

```go
func StampaInfo(valore int, puntatore *int)
```
Stampa il valore della variabile, l'indirizzo di memoria (usando `%p`), e il valore dereferenziato dal puntatore.

## Vincoli
- Nella funzione `main`, crea due variabili intere e stampane gli indirizzi
- Chiama `Scambia` per scambiare i valori
- Verifica che i valori siano effettivamente scambiati dopo la chiamata

## Esempio

Output:
```
Prima dello scambio: x = 5, y = 10
Indirizzo di x: 0xc000016098
Indirizzo di y: 0xc0000160a0
Dopo lo scambio: x = 10, y = 5
```

## Suggerimento
- In Go, `&variabile` restituisce l'indirizzo di memoria
- `*puntatore` accede al valore puntato (dereferenziazione)
- I parametri delle funzioni sono passati per valore: senza puntatori, le modifiche non sopravvivono alla funzione
