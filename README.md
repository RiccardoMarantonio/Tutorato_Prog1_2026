# Tutorato Programmazione 1 — Laboratorio Go

Repository didattica per il corso di Programmazione 1. Contiene **21 esercizi** organizzati in **6 capitoli** tematici, ciascuno con 4 livelli di difficoltà crescente: **Intro**, **Easy**, **Medium**, **Hard**.

Ogni esercizio è autonomo: ha una traccia (`README.md`), uno skeletro da completare (`main.go`) e test automatici (`main_test.go`).

## Prerequisiti

- Go 1.22+
- Un editor di testo o IDE con supporto Go
- Terminale

## Struttura della Repository

```
├── exercises/                          # Esercizi per gli studenti
│   ├── 01_if_for_switch_functions/     # Capitolo 1: controllo flusso e funzioni
│   │   ├── 01_numeri_pari/             #   Intro
│   │   ├── 02_indovina_numero/         #   Easy
│   │   ├── 03_fizzbuzz_avanzato/       #   Medium
│   │   └── 04_triangolo_numeri/        #   Hard
│   ├── 02_array_slice_string_rune/     # Capitolo 2: array, slice, stringhe
│   │   ├── 01_somma_array/
│   │   ├── 02_max_min_slice/
│   │   ├── 03_inversione_palindromi/
│   │   └── 04_fusione_slice_ordinate/
│   ├── 03_puntatori_type_struct/       # Capitolo 3: puntatori, type, struct
│   │   ├── 01_intro_puntatori/
│   │   ├── 02_rettangolo_struct/
│   │   ├── 03_rubrica_telefonica/
│   │   └── 04_gestione_magazzino/
│   ├── 04_mappe/                       # Capitolo 4: mappe
│   │   ├── 01_contatore_parole/
│   │   ├── 02_anagrammi/
│   │   ├── 03_voti_studenti/
│   │   └── 04_elezioni/
│   ├── 05_ricorsione/                  # Capitolo 5: ricorsione
│   │   ├── 01_somma_ricorsiva/
│   │   ├── 02_fibonacci/
│   │   ├── 03_fibonacci_hanoi/
│   │   └── 04_espressioni_aritmetiche/
│   └── 06_mini_progetto_ripasso/       # Capitolo 6: progetto finale
│       └── 01_gestione_biblioteca/
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
- Descrizione del problema
- Firme delle funzioni da implementare
- Vincoli e restrizioni
- Esempi di input/output
- Suggerimenti

### 2. Completa lo skeletro

Apri il file `main.go` dell'esercizio. Troverai le funzioni con il corpo `// TODO: implementare` e un `return` placeholder. Sostituisci con la tua implementazione.

### 3. Esegui i test

```bash
# Test di un singolo esercizio
go test ./exercises/01_if_for_switch_functions/01_numeri_pari

# Test con output dettagliato (vedi ogni sub-test)
go test -v ./exercises/01_if_for_switch_functions/01_numeri_pari

# Con Make
make test-exercise EX=01_if_for_switch_functions/01_numeri_pari
make test-exercise-v EX=01_if_for_switch_functions/01_numeri_pari
```

### 4. Esegui il programma

```bash
# Esegui l'esercizio (per testare l'I/O da stdin)
go run ./exercises/01_if_for_switch_functions/01_numeri_pari

# Con Make
make run EX=01_if_for_switch_functions/01_numeri_pari
```

### Output dei Test

**Quando i test passano:**
```
=== RUN   TestIsEven
=== RUN   TestIsEven/zero
=== RUN   TestIsEven/due
--- PASS: TestIsEven (0.00s)
    --- PASS: TestIsEven/zero (0.00s)
    --- PASS: TestIsEven/due (0.00s)
PASS
```

**Quando i test falliscono:**
```
--- FAIL: TestIsEven (0.00s)
    --- FAIL: TestIsEven/zero (0.00s)
        main_test.go:21: IsEven(0) = false, want true
FAIL
```
Il messaggio ti dice: file, riga, valore ottenuto (`got`) e valore atteso (`want`).

## Catalogo Esercizi

### Capitolo 1 — If, For, Switch, Funzioni
*Concetti: tipi base, if/else, for, switch, funzioni, return*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Numeri Pari | Intro | Stampa i numeri pari da 0 a n con funzione `IsEven` |
| 02 | Indovina il Numero | Easy | Gioco interattivo: indovina 42 con feedback alto/basso |
| 03 | FizzBuzz Avanzato | Medium | FizzBuzz + numeri primi, con switch e funzioni helper |
| 04 | Triangolo di Numeri | Hard | Stampa triangolo di numeri consecutivi allineati |

### Capitolo 2 — Array, Slice, Stringhe, Rune
*Concetti: array a dimensione fissa, slice dinamiche, stringhe come byte, rune*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Somma di un Array | Intro | Somma 5 numeri in un array `[5]int` |
| 02 | Max/Min Slice | Easy | Trova massimo e minimo in una slice |
| 03 | Inversione e Palindromi | Medium | Inverte slice, verifica palindromi stringa e numerici |
| 04 | Fusione di Slice Ordinate | Hard | Merge efficiente, rimuovi duplicati, intersezione |

### Capitolo 3 — Puntatori, Type, Struct
*Concetti: puntatori, dereferenziazione, type personalizzati, struct, passaggio per valore/riferimento*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Intro ai Puntatori | Intro | Scambia due variabili con puntatori, stampa indirizzi |
| 02 | Rettangolo con Struct | Easy | Struct `Rettangolo` con Area, Perimetro, Scala |
| 03 | Rubrica Telefonica | Medium | CRUD contatti con struct, ricerca case-insensitive |
| 04 | Gestione Magazzino | Hard | Prodotti con carico/scarico, validazioni, statistiche |

### Capitolo 4 — Mappe
*Concetti: mappe, chiavi/valori, iterazione, mappe con struct*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Contatore di Parole | Intro | Conta occorrenze parole con `map[string]int` |
| 02 | Anagrammi | Easy | Raggruppa anagrammi tramite firma (caratteri ordinati) |
| 03 | Voti degli Studenti | Medium | Registro voti con mappe, medie, distribuzioni |
| 04 | Elezioni | Hard | Sezioni elettorali, risultati, affluenza, ballottaggio |

### Capitolo 5 — Ricorsione
*Concetti: caso base, passo ricorsivo, stack di chiamate, memoizzazione*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Somma Ricorsiva | Intro | Somma da 1 a n con ricorsione semplice |
| 02 | Fibonacci | Easy | Fibonacci ricorsivo + sequenza completa |
| 03 | Fibonacci + Hanoi | Medium | Fibonacci con memoizzazione + Torre di Hanoi |
| 04 | Espressioni Aritmetiche | Hard | Parser ricorsivo per espressioni con parentesi |

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
make test-exercise EX=03_puntatori_type_struct/02_rettangolo_struct
make run EX=03_puntatori_type_struct/02_rettangolo_struct
```

## Progressione Consigliata

1. **Segui l'ordine dei capitoli** — ogni capitolo introduce concetti nuovi che servono per il successivo
2. **Dentro ogni capitolo, procedi in ordine** — Intro → Easy → Medium → Hard
3. **Non saltare l'Intro** — serve a familiarizzare con il formato e i concetti base
4. **Se un esercizio è troppo difficile** — torna al precedente, consolida le basi
5. **Il progetto finale (capitolo 6)** va fatto dopo aver completato almeno i capitoli 1-5

## Cosa NON Usare

Gli esercizi sono progettati per essere risolti con **solo le librerie standard di base** (`fmt`, etc...). Non usare:
- `sort` package (ordina con cicli)
- `strings` package (implementa le utility a mano)
- `strconv` (tranne dove esplicitamente permesso)
- Package esterni di qualsiasi tipo
