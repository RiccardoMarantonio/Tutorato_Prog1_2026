# Suggerimenti Generali - Capitolo 4: mappe

Ecco alcuni suggerimenti generali sulla sintassi e le funzionalità di Go che potrebbero esserti utili per risolvere gli esercizi di questo capitolo:

- **Dichiarazione e inizializzazione di mappe**: Una mappa in Go deve essere esplicitamente inizializzata prima di potervi aggiungere degli elementi, altrimenti si verificherà un *panic*. Il modo idiomatico per farlo è: `miaMappa := make(map[string]int)`.
- **Valori di default (Zero-values)**: Se si tenta di leggere da una mappa usando una chiave che non esiste, Go restituisce comodamente lo *zero-value* del tipo associato ai valori (es. `0` per gli `int`, `""` per le `string`). Questo consente di incrementare direttamente contatori (`mappa[chiave]++`) senza prima dover verificare se la chiave esista.
- **Verificare la presenza di una chiave**: Se è necessario distinguere tra una chiave assente e una chiave presente ma il cui valore è lo zero-value, è possibile sfruttare il prelievo a due valori: `valore, esiste := mappa[chiave]`. `esiste` sarà un booleano.
- **Mappe di Puntatori**: L'utilizzo di valori puntatore all'interno della mappa (es. `map[string]*MioTipo`) è molto utile qualora si debbano aggiornare frequentemente e direttamente i campi delle struct interne (es. `mappa["chiave"].CampoPuntatore = nuovoValore`) senza dover riassegnare l'intera struct alla mappa.
- **Iterazione sulle mappe**: Si può iterare in modo semplice su tutte le coppie chiave-valore di una mappa usando: `for chiave, valore := range mappa { ... }`. Ricorda che l'ordine di iterazione nelle mappe in Go è volutamente casuale e non garantito.
- **Lettura dell'input**: Il pacchetto `bufio` (e specificatamente `bufio.NewScanner(os.Stdin)`) è uno strumento eccellente ed efficiente per leggere input complessi riga per riga. In combinazione, la funzione `strings.Fields(riga)` è molto comoda per dividere in automatico una riga in una slice di parole basandosi sugli spazi.
