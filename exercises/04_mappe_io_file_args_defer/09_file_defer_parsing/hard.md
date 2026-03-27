# File/Defer/Parsing - Hard - Migliore Media Da File

Implementa BestAverage(path string) (string, float64, error).

Formato file:
- una riga: nome,voto

Regole:
- voto valido nel range 0..30
- ignora righe con voto non valido
- seleziona lo studente con media piu alta
- in parita scegli nome lessicograficamente minore
- errore se nessuno studente ha almeno un voto valido