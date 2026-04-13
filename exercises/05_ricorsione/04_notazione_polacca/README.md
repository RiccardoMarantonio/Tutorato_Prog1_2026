# Interprete in Notazione Polacca (Prefissa)

Questo esercizio introduce la valutazione di espressioni aritmetiche utilizzando la **Notazione Polacca** (Polish Notation, PN), conosciuta anche come notazione prefissa.

Il programma legge da standard input una riga contenente un'espressione in notazione prefissa, separata da spazi, e ne stampa il risultato. In caso di errore, stampa un messaggio di errore.

Il programma deve implementare le seguenti funzioni:

```go
func ValutaPN(tokens []string) (int, error)
```
Valuta l'intera espressione in notazione prefissa a partire da uno slice di stringhe (estratto dall'input letto).

```go
func valutaRicorsivo(tokens []string) (int, []string, error)
```
Funzione helper ricorsiva che estrae l'operando corrente o calcola ricorsivamente il risultato degli operatori, restituendo il valore calcolato, i restanti token non ancora consumati, e un eventuale errore.

## Introduzione alla Notazione Polacca

Normalmente siamo abituati a scrivere le espressioni aritmetiche in **notazione infissa**, dove l'operatore si trova *tra* i due operandi:
`3 + 4`

Nella **notazione prefissa (PN)**, l'operatore viene scritto *prima* dei suoi operandi:
`+ 3 4`

Come per la notazione postfissa, il grande vantaggio della PN è che **non richiede mai l'uso delle parentesi** per stabilire la precedenza delle operazioni! L'ordine di esecuzione è implicitamente determinato dalla posizione degli operatori e dalla loro "arità" (cioè quanti operandi richiedono, in questo caso 2).


## Vincoli
- Il programma deve leggere l'espressione da **standard input** (es. con `bufio.Scanner`).
- i token sono suddivisi da spazi
- Supporta le operazioni: `+`, `-`, `*`, `/` (divisione intera)
- Tutti i numeri in input saranno validi interi convertibili (es. `"10"`, `"-5"`).
- Se la divisione comporta un divisore uguale a zero, ritorna un errore.
- Se alla fine dell'espressione rimangono token non consumati, o se un operatore richiede operandi mancanti, la funzione deve restituire un errore.
- Il programma stamperà `Risultato: <n>` se va a buon fine, o `Errore: <messaggio>` in caso di errore.

## Esempio

Input:
```
- + 5 * 1 2 3
```

Output:
```
Risultato: 4
```

## Suggerimento
- Per la lettura da stdin, puoi usare uno scanner (`bufio.NewScanner(os.Stdin)`) per leggere la riga e poi `strings.Fields(testo)` per ottenere direttamente lo slice di stringhe eliminando gli spazi extra.
- Caso base: se il token è un numero (convertibile con `strconv.Atoi`), ritornalo con i restanti token.
- Passo ricorsivo: se il token è un operatore, chiama `valutaRicorsivo` sui token rimanenti per calcolare il primo operando, e di nuovo sui token restanti dopo aver calcolato il primo operando per calcolare il secondo.
