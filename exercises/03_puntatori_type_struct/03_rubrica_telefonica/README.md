# Rubrica Telefonica Semplice

Crea un sistema di gestione di una rubrica telefonica usando struct.

Definisci le seguenti strutture:

```go
type Contatto struct {
	Nome     string
	Cognome  string
	Telefono string
}

type Rubrica struct {
	Contatti []Contatto
}
```

Implementa le seguenti funzioni:

```go
func Aggiungi(rubrica *Rubrica, c Contatto)
```
Aggiunge un contatto alla rubrica. Se un contatto con lo stesso numero di telefono esiste già, non aggiungerlo.

```go
func CercaPerNome(rubrica Rubrica, nome string) []Contatto
```
Restituisce tutti i contatti che hanno il nome specificato (confronto case-insensitive).

```go
func Rimuovi(rubrica *Rubrica, telefono string) bool
```
Rimuove il contatto con il numero di telefono specificato. Restituisce `true` se trovato e rimosso, `false` altrimenti.

```go
func StampaRubrica(rubrica Rubrica)
```
Stampa tutti i contatti nella rubrica in formato: `Cognome, Nome - Telefono`.

## Vincoli
- La rubrica deve essere passata per puntatore a `Aggiungi` e `Rimuovi` (per modificarla)
- La ricerca può ricevere la rubrica per valore (non la modifica)
- Per il confronto case-insensitive, converti tutto in minuscolo confrontando byte per byte

## Esempio

Input:
```
AGGIUNGI Mario Rossi 3331234567
AGGIUNGI Luigi Bianchi 3339876543
AGGIUNGI Mario Verdi 3331112222
CERCA Mario
STAMPA
RIMUOVI 3339876543
STAMPA
FINE
```

Output:
```
Contatto aggiunto: Mario Rossi
Contatto aggiunto: Luigi Bianchi
Contatto aggiunto: Mario Verdi

Risultati ricerca "Mario":
Rossi, Mario - 3331234567
Verdi, Mario - 3331112222

=== Rubrica ===
Rossi, Mario - 3331234567
Bianchi, Luigi - 3339876543
Verdi, Mario - 3331112222

Contatto rimosso: 3339876543

=== Rubrica ===
Rossi, Mario - 3331234567
Verdi, Mario - 3331112222
```
