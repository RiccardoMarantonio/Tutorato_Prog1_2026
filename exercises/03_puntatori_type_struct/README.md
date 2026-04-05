# Suggerimenti Generali - Capitolo 3: puntatori, type, struct

Ecco alcuni suggerimenti generali sulla sintassi e le funzionalità di Go che potrebbero esserti utili per risolvere gli esercizi di questo capitolo:

- **Indirizzi di memoria**: Usa `&variabile` per ottenere l'indirizzo di memoria (puntatore) di una variabile.
- **Dereferenziazione**: Usa `*puntatore` per accedere o modificare il valore salvato all'indirizzo di memoria a cui il puntatore fa riferimento.
- **Passaggio di parametri**: Passare una `struct` a una funzione per *valore* ne crea una copia (le modifiche non avranno effetto sull'originale); passarla per *puntatore* (es. `*MioTipo`) permette di modificarne i campi.
- **Accesso ai campi di una struct**: Si accede ai campi usando la notazione con il punto (es. `p.X`). In Go, il dereferenziamento dei puntatori a struct avviene automaticamente, quindi si usa `r.Base` sia che `r` sia un valore sia che sia un puntatore.
- **Rimozione da una slice**: Per rimuovere un elemento all'indice `i` da una slice, un pattern efficace in Go è unire gli elementi precedenti e successivi usando `append(slice[:i], slice[i+1:]...)`.
- **Gestione dell'input**: Per interpretare comandi testuali in sequenza, funzioni come `fmt.Scan` o `fmt.Scanln` risultano molto comode per consumare l'input parola per parola o riga per riga.
- **Gestione degli errori**: Puoi creare un nuovo errore formattato da restituire nelle tue funzioni usando `fmt.Errorf("messaggio di errore")`.
