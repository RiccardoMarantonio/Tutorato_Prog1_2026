# I/O e Argomenti - Hard - Mini Interprete

Implementa RunProgram(r io.Reader) (int, error).

Lettura riga per riga di comandi:
- set N
- add N
- sub N
- mul N
- end

Regole:
- stato iniziale 0
- end termina il programma
- comando non valido => errore
- se manca end => errore