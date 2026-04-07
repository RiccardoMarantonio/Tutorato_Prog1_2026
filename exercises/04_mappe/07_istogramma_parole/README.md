# Istogramma Parole

Scrivi un programma che legge un file di testo e conta la frequenza di ogni parola, poi stampa un istogramma a barre orizzontali ordinato dalla parola più frequente alla meno frequente.

## Requisiti

1. Leggi il percorso del file da riga di comando (`os.Args[1]`)
2. Leggi tutto il contenuto del file
3. Suddividi il testo in parole (usa `strings.Fields` per splittare)
4. Conta le occorrenze di ogni parola usando una mappa
5. Ordina le parole per frequenza decrescente
6. Stampa l'istogramma con barre orizzontali usando il carattere `#`

## Formato Output

```
Parola più frequente  : ########## 10
Seconda parola       : ########   8
Terza parola         : ######     6
```

- Larghezza massima della barra: 50 caratteri `#`
- Allineamento a destra per le frequenze numeriche
- Stampa solo le parole con frequenza >= 2

## Esempio

Input (file `testo.txt`):
```
il gatto mangia il pesce il gatto dorme
il cane corre il gatto miagola
```

Output:
```
il          : ################ 4
gatto       : ############### 3
```

## Funzioni da implementare

```go
func LeggiFile(path string) (string, error)
func EstraiParole(testo string) []string
func ContaFrequenze(parole []string) map[string]int
func OrdinaPerFrequenza(frequenze map[string]int) []Pair
func StampaIstogramma(paroleOrdinate []Pair)
```

Dove `Pair` e una struct con `Parola string` e `Frequenza int`.

## Vincoli
- Ignora la distinzione tra maiuscole e minuscole (trasforma tutto in minuscolo)
- Rimuovi la punteggiatura dalle parole
- Usa le funzioni di `strings` e `sort`
- NON usare funzioni esterne o librerie non standard
