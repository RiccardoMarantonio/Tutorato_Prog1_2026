# Cifrario di Cesare

Implementa il **cifrario di Cesare**, una tecnica di cifratura per sostituzione in cui ogni lettera del testo in chiaro viene shiftata di un numero fisso di posizioni nell'alfabeto.

Ad esempio, con chiave `k = 3`:
- `d` diventa `a`
- `e` diventa `b`
- `c` diventa `z` (wrap-around)

Il programma deve leggere la chiave `k` dalla **prima riga** e il testo cifrato dalle **righe successive** (input multilinea). Il testo decifrato va stampato su stdout.

Il programma deve implementare la funzione:

```go
func Decrypt(s string, k int) string
```

## Esempi

Input:
```
3
fldr
```

Output:
```
Decifrato: ciao
```

Input:
```
3
def ghil
abc xyz
```

Output:
```
Decifrato: abc def
xyz uvw
```

Input:
```
1
bcd
```

Output:
```
Decifrato: abc
```

## Vincoli
- Solo lettere minuscole ASCII (a-z), spazi e newline
- La chiave `k` e un intero non negativo
- I caratteri non alfabetici (spazi, newline) devono rimanere invariati
- Il testo cifrato puo occupare piu righe

## Suggerimento
- Per decifrare: `(char - 'a' - k%26 + 26) % 26 + 'a'`
- Usa `[]rune` per iterare sulla stringa
- Leggi tutto il resto dell'input dopo la prima riga con `bufio.Scanner`
