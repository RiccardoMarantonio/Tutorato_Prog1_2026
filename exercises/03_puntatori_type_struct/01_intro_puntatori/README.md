# Introduzione a Puntatori e Struct

Scrivi un programma che dimostra il funzionamento base dei puntatori e delle struct in Go.

## Parte 1: Puntatori

Implementa la funzione:

```go
func Scambia(a, b *int)
```
Scambia i valori di due variabili intere usando i puntatori.

## Parte 2: Struct

Definisci una struct `Punto` per rappresentare un punto nel piano cartesiano:

```go
type Punto struct {
    X float64
    Y float64
}
```

Implementa le seguenti funzioni:

```go
func DistanzaDaOrigine(p Punto) float64
```
Calcola la distanza del punto dall'origine (0, 0). Formula: `sqrt(X*X + Y*Y)`.

```go
func Sposta(p *Punto, dx, dy float64)
```
Sposta il punto di `dx` sull'asse X e `dy` sull'asse Y. Usa un puntatore per modificare la struct originale.

## Esempio

Input:
```
3 4
1 1
```

Output:
```
=== Puntatori ===
Prima: x = 5, y = 10
Dopo:  x = 10, y = 5

=== Struct Punto ===
Punto: (3.00, 4.00)
Distanza dall'origine: 5.00
Dopo spostamento (1.00, 1.00): (4.00, 5.00)
```

## Vincoli
- La funzione `Scambia` deve usare i puntatori
- La funzione `Sposta` deve modificare la struct originale (passa per puntatore)
- `DistanzaDaOrigine` puo ricevere la struct per valore
- Per la radice quadrata, implementa `Sqrt(x float64) float64` con il metodo di Newton o usa un'approssimazione
