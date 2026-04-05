# Suggerimenti Generali - Capitolo 5: ricorsione

Ecco alcuni suggerimenti generali sulla sintassi e le funzionalità di Go (e concetti architetturali) che potrebbero esserti utili per risolvere gli esercizi di questo capitolo:

- **Struttura della Ricorsione**: Ogni funzione ricorsiva ha bisogno di almeno un *caso base* che fermi l'esecuzione per non incorrere in un loop infinito (o uno stack overflow), e di un *passo ricorsivo* in cui la funzione chiama sé stessa con parametri che si avvicinano progressivamente al caso base.
- **Mappe e Slice come cache (Memoizzazione)**: Dato che in Go `map` e `slice` agiscono come reference-type per la struttura dati interiore, passarli come parametri alle successive chiamate ricorsive permette di leggere e modificare i dati condivisi (ad esempio per implementare una "memoria" o cache e salvare le risposte intermedie) senza dover usare esplicitamente puntatori aggiuntivi.
- **Funzioni Helper/Wrapper**: Un pattern diffusissimo in Go è creare una funzione primaria che riceve in input solo ciò che serve all'utente, la quale inizializza le strutture ausiliarie (come le mappe di cache) e invoca al suo interno una funzione ricorsiva *helper*, che riceve a sua volta la cache e le altre variabili di supporto.
