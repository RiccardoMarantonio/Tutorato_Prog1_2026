# Rettangolo con Struct

Definisci un tipo `Rettangolo` come struct con i campi `Base` e `Altezza` (entrambi `float64`).

Il programma deve implementare le seguenti funzioni:

```go
func Area(r Rettangolo) float64
```
Restituisce l'area del rettangolo.

```go
func Perimetro(r Rettangolo) float64
```
Restituisce il perimetro del rettangolo.

```go
func Scala(r *Rettangolo, fattore float64)
```
Modifica il rettangolo originale moltiplicando base e altezza per il fattore dato. Usa un puntatore per modificare la struct originale.

## Vincoli
- Definisci il tipo con `type Rettangolo struct { ... }`
- La funzione `Scala` deve modificare la struct originale (passa per puntatore)
- Leggi base e altezza da stdin

## Esempio

Input:
```
4 6
2
```

Output:
```
Rettangolo: Base=4.00, Altezza=6.00
Area: 24.00
Perimetro: 20.00
Dopo scala x2: Base=8.00, Altezza=12.00
Nuova area: 96.00
```

## Suggerimento
- Passare una struct per valore crea una copia
- Passare una struct per puntatore (`*Rettangolo`) permette di modificarla
- Accedi ai campi con `r.Base` sia per valore che per puntatore (Go fa il dereferenziamento automatico)
