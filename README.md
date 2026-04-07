# Tutorato Programmazione 1 — Laboratorio Go

Repository didattica per il corso di Programmazione 1. Contiene **35 esercizi** organizzati in **6 capitoli** tematici, con difficoltà crescente.

Ogni esercizio è autonomo: ha una traccia (`README.md`), uno scheletro da completare (`main.go`) e test automatici (`main_test.go`).

## Prerequisiti

- Go 1.22+
- Un editor di testo o IDE con supporto Go
- Terminale

## Struttura della Repository

```
├── exercises/                          # Esercizi per gli studenti
│   ├── 01_if_for_switch_functions/     # Capitolo 1: controllo flusso e funzioni
│   ├── 02_array_slice_string_rune/     # Capitolo 2: array, slice, stringhe, rune
│   ├── 03_puntatori_type_struct/       # Capitolo 3: puntatori, type, struct
│   ├── 04_mappe/                       # Capitolo 4: mappe
│   ├── 05_ricorsione/                  # Capitolo 5: ricorsione
│   └── 06_mini_progetto_ripasso/       # Capitolo 6: progetto finale
├── solutions/                          # Soluzioni complete (per tutor)
│   └── ... (stessa struttura di exercises/)
├── test/
│   └── helpers.go                      # Utility comuni per i test
├── Makefile
├── go.mod
└── README.md
```

## Come Lavorare su un Esercizio

### 1. Leggi la traccia

Ogni esercizio ha un `README.md` con:
- Descrizione del problema in linguaggio naturale
- Strutture dati da definire
- Funzionalit da implementare
- Vincoli e restrizioni
- Esempi di input/output

### 2. Completa lo scheletro

Apri il file `main.go` dell'esercizio. Troverai funzioni con `// TODO: implementa` o un `return` placeholder. In alcuni esercizi il `main()` e gia fornito, in altri devi scriverlo tu.

### 3. Esegui i test

```bash
# Test di un singolo esercizio
go test ./exercises/04_mappe/07_mappe_comme_set

# Test con output dettagliato
go test -v ./exercises/04_mappe/07_mappe_comme_set

# Con Make
make test-exercise EX=04_mappe/07_mappe_comme_set
make test-exercise-v EX=04_mappe/07_mappe_comme_set
```

### 4. Esegui il programma

```bash
go run ./exercises/04_mappe/07_mappe_comme_set

# Con Make
make run EX=04_mappe/07_mappe_comme_set
```

### Output dei Test

**Quando i test passano:**
```
--- PASS: TestSetBasics (0.00s)
```

**Quando i test falliscono:**
```
--- FAIL: TestSetBasics (0.00s)
    main_test.go:10: dopo Add, Contains dovrebbe restituire true
```
Il messaggio ti dice quale asserzione e fallita.

## Catalogo Esercizi

### Capitolo 1 — If, For, Switch, Funzioni
*Concetti: tipi base, if/else, for, switch, funzioni, return, stringhe*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Numeri Pari | Intro | Stampa i numeri pari da 0 a n |
| 02 | Indovina il Numero | Easy | Gioco: indovina 42 con feedback |
| 03 | Calcolatrice | Easy | Calcolatrice con switch |
| 04 | Indovina la Parola | Easy | Indovina parola con 6 tentativi |
| 05 | FizzBuzz Avanzato | Medium | FizzBuzz + numeri primi con switch |
| 06 | Tabelline | Medium | Stampa tabelline con cicli annidati |
| 07 | Numeri Perfetti | Medium | Trova numeri perfetti fino a n |
| 08 | Triangolo di Numeri | Hard | Triangolo di numeri allineati |
| 09 | Pattern Stelle | Hard | Triangolo, diamante, piramide numerica |

### Capitolo 2 — Array, Slice, Stringhe, Rune
*Concetti: array fissi, slice dinamiche, stringhe come byte, rune Unicode*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Somma di un Array | Intro | Somma 5 numeri in array fisso |
| 02 | Max/Min Slice | Intro | Trova massimo e minimo in slice |
| 03 | Medie di un Array | Easy | Media, conta sopra/sotto media |
| 04 | Rotazione Slice | Easy | Ruota slice di k posizioni |
| 05 | Palindromi Stringhe | Easy | Verifica palindromi case-insensitive |
| 06 | Conteggio Rune | Easy | Conta lettere, cifre, spazi, altri |
| 07 | Rotazione Stringhe | Medium | Ruota caratteri in stringa |
| 08 | Cifrario di Cesare | Medium | Cifra/decifra con shift |
| 09 | Anagrammi Rune | Medium | Verifica anagrammi Unicode |
| 10 | Fusione Slice Ordinate | Medium | Merge, rimuovi duplicati, intersezione |
| 11 | Wordle | Hard | Gioco Wordle in console |

### Capitolo 3 — Puntatori, Type, Struct
*Concetti: puntatori, dereferenziazione, type personalizzati, struct*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Intro ai Puntatori | Intro | Scambia variabili con puntatori |
| 02 | Rettangolo con Struct | Easy | Struct con Area, Perimetro, Scala |
| 03 | Rubrica Telefonica | Medium | CRUD contatti, ricerca case-insensitive |
| 04 | Gestione Magazzino | Hard | Prodotti con carico/scarico, statistiche |

### Capitolo 4 — Mappe
*Concetti: mappe, chiavi/valori, iterazione, set*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Contatore di Parole | Intro | Conta occorrenze con `map[string]int` |
| 02 | Anagrammi | Easy | Raggruppa anagrammi tramite firma |
| 03 | Gestione Magazzino | Medium | Prodotti con map, statistiche |
| 04 | Voti degli Studenti | Medium | Registro voti, medie, distribuzioni |
| 05 | Mappe come Set | Medium | Implementa Set con `map[string]struct{}` |
| 06 | Elezioni | Hard | Sezioni elettorali, affluenza, ballottaggio |
| 07 | Istogramma Parole | Hard | Leggi file, conta frequenze, stampa istogramma |

### Capitolo 5 — Ricorsione
*Concetti: caso base, passo ricorsivo, stack di chiamate, memoizzazione*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Somma Ricorsiva | Intro | Somma da 1 a n con ricorsione |
| 02 | Fibonacci con Memo | Easy | Fibonacci ottimizzato + conta chiamate |
| 03 | Torre di Hanoi | Medium | Risolutore ricorsivo della Torre di Hanoi |
| 04 | Espressioni Aritmetiche | Hard | Parser ricorsivo per espressioni |

### Capitolo 6 — Mini Progetto Ripasso
*Concetti: tutti i precedenti combinati*

| # | Esercizio | Cosa si fa |
|---|-----------|------------|
| 01 | Gestione Biblioteca | Sistema completo: libri, prestiti, ricerche, statistiche |

## Comandi Make

| Comando | Descrizione |
|---------|-------------|
| `make test-exercise EX=<percorso>` | Test di un esercizio |
| `make test-exercise-v EX=<percorso>` | Test verbose (dettagliato) |
| `make run EX=<percorso>` | Esegue un esercizio |
| `make test` | Esegue tutti i test della repository |
| `make fmt` | Formatta tutto il codice con gofmt |

Esempio:
```bash
make test-exercise EX=04_mappe/07_mappe_comme_set
make run EX=04_mappe/07_mappe_comme_set
```

## Progressione Consigliata

1. **Segui l'ordine dei capitoli** — ogni capitolo introduce concetti nuovi che servono per il successivo
2. **Dentro ogni capitolo, procedi in ordine** — dal più semplice al più difficile
3. **Non saltare l'Intro** — serve a familiarizzare con il formato e i concetti base
4. **Se un esercizio e troppo difficile** — torna al precedente, consolida le basi
5. **Il progetto finale (capitolo 6)** va fatto dopo aver completato almeno i capitoli 1-5

## Cosa NON Usare

Gli esercizi sono progettati per essere risolti con **solo le librerie standard di base** (`fmt`, `bufio`, `os`). Non usare:
- `sort` package (ordina con cicli)
- `strings` package (implementa le utility a mano)
- `strconv` (tranne dove esplicitamente permesso)
- Package esterni di qualsiasi tipo

## Note per i Tutor

- Le soluzioni complete sono in `solutions/` con la stessa struttura di `exercises/`
- I test sono **solo** nella cartella `exercises/` — le soluzioni non hanno test duplicati
- `test/helpers.go` contiene utility comuni per testare I/O da stdin
- Per verificare una soluzione: `go build ./solutions/<capitolo>/<esercizio>`
