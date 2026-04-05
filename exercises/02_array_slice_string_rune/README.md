# Suggerimenti Generali - Capitolo 2: array, slice, string, rune

Ecco alcuni suggerimenti generali sulla sintassi e le funzionalità di Go che potrebbero esserti utili per risolvere gli esercizi di questo capitolo:

- **Dimensione fissa degli array**: Un array in Go ha sempre una dimensione fissa definita a compile-time, specificata tra parentesi quadre (es. `[5]int`).
- **Creare slice**: Per creare una nuova slice usa la funzione integrata `make`, specificando il tipo e la lunghezza: `make([]int, len(slice))`.
- **Stringhe e Rune (Unicode)**: Per lavorare correttamente con i caratteri che includono encoding multibyte (come accenti o emoji), converti la stringa in una slice di rune usando `runes := []rune(s)`. Per fare l'operazione opposta, usa `s = string(runes)`.
- **Iterare sulle strighe**: Usando `for i, r := range s`, il ciclo itera automaticamente sulle rune (caratteri Unicode) e non sui singoli byte della stringa.
- **Controllo di range di caratteri ASCII**: Puoi facilmente verificare se una runa rappresenta una lettera o una cifra confrontandola direttamente con i caratteri scritti tra singoli apici: `r >= 'a' && r <= 'z'`, `r >= 'A' && r <= 'Z'` per le lettere, `r >= '0' && r <= '9'` per le cifre. Oppure utilizzando le apposite funzioni `isDigit`, `isLetter`, etc del package `unicode`.