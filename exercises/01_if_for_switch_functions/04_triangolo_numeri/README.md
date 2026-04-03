# Triangolo di Numeri

Scrivi un programma che, dato un numero `n` da stdin, stampa un triangolo di numeri con le seguenti proprietà:

- Il triangolo ha `n` righe
- Ogni riga `i` (partendo da 1) contiene `i` numeri
- I numeri sono consecutivi a partire da 1
- Ogni numero deve essere stampato allineato a destra in un campo di larghezza 4 (usa `fmt.Printf("%4d", num)`)

Il programma deve implementare le seguenti funzioni:

```go
func CalcolaNumeroRiga(riga int) int
```
Restituisce l'ultimo numero presente nella riga specificata (es: riga 3 → 6, perché i numeri sono 1,2,3,4,5,6).

```go
func SommaTriangolo(n int) int
```
Restituisce la somma di tutti i numeri nel triangolo di `n` righe.

```go
func StampaTriangolo(n int)
```
Stampa il triangolo formattato.

## Vincoli
- `1 <= n <= 30`
- Allineamento obbligatorio con `fmt.Printf`
- La funzione `StampaTriangolo` non deve restituire nulla, solo stampare

## Esempio

Input:
```
5
```

Output:
```
   1
   2   3
   4   5   6
   7   8   9  10
  11  12  13  14  15
```

Output aggiuntivo (dopo il triangolo):
```
Ultimo numero riga 5: 15
Somma totale: 120
```

## Estensione (opzionale)
Prova a stampare il triangolo in forma di piramide centrata.
