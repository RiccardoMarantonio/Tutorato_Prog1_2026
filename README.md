# Tutorato Programmazione 1 — Laboratorio Go

Repository didattica per il corso di Programmazione 1. Contiene **28 esercizi** organizzati in **6 capitoli** tematici, con difficoltà crescente da **Intro** a **Hard**.

Ogni esercizio è autonomo: ha una traccia (`README.md`), uno scheletro da completare (`main.go`) e test automatici (`main_test.go`).

## Prerequisiti

- Go 1.22+
- Un editor di testo o IDE con supporto Go
- Terminale

## Struttura della Repository

```
├── exercises/                          # Esercizi per gli studenti
│   ├── 01_if_for_switch_functions/     # Capitolo 1: controllo flusso e funzioni (7 esercizi)
│   │   ├── 01_numeri_pari/             #   Intro — solo funzione, main da scrivere
│   │   ├── 02_indovina_numero/         #   Easy — funzione + main fornito
│   │   ├── 03_fizzbuzz_avanzato/       #   Medium — 3 funzioni da implementare
│   │   ├── 04_triangolo_numeri/        #   Hard — 3 funzioni + output formattato
│   │   ├── 05_calcolatrice/            #   Extra — switch, scrivi tutto da zero
│   │   ├── 06_tabelline/               #   Extra — cicli annidati, main da scrivere
│   │   └── 07_indovina_parola/         #   Extra — stringhe + len(), main da scrivere
│   ├── 02_array_slice_string_rune/     # Capitolo 2: array, slice, stringhe, rune (8 esercizi)
│   │   ├── 01_somma_array/             #   Array — somma array fisso, main da scrivere
│   │   ├── 02_medie_array/             #   Array — media, sopra/sotto media
│   │   ├── 03_max_min_slice/           #   Slice — max/min, main da scrivere
│   │   ├── 04_rotazione_slice/         #   Slice — rotazione, main da scrivere
│   │   ├── 05_palindromi_stringhe/     #   Stringhe — palindromi case-insensitive
│   │   ├── 06_conta_rune/              #   Rune — conteggio Unicode, main da scrivere
│   │   ├── 07_anagrammi_rune/          #   Rune — anagrammi Unicode, main da scrivere
│   │   └── 08_fusione_slice_ordinate/  #   Slice — merge, intersezione, main da scrivere
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
│   │   ├── 02_fibonacci_memo/
│   │   ├── 03_torre_hanoi/
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
- Firme delle funzioni da implementare (dove applicabile)
- Vincoli e restrizioni
- Esempi di input/output
- Suggerimenti

### 2. Completa lo scheletro

Apri il file `main.go` dell'esercizio. Troverai funzioni con `// TODO: implementare` e un `return` placeholder. In alcuni esercizi il `main()` e gia fornito, in altri devi scriverlo tu (seguito da commenti guida).

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
*Concetti: tipi base, if/else, for, switch, funzioni, return, stringhe (len)*

| # | Esercizio | Formato | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Numeri Pari | Funzione + main da scrivere | Stampa i numeri pari da 0 a n |
| 02 | Indovina il Numero | Funzioni + main fornito | Gioco: indovina 42 con feedback |
| 03 | FizzBuzz Avanzato | 3 funzioni + main fornito | FizzBuzz + numeri primi con switch |
| 04 | Triangolo di Numeri | 3 funzioni + main fornito | Triangolo di numeri allineati |
| 05 | Calcolatrice | Funzione + main da scrivere | Calcolatrice con switch |
| 06 | Tabelline | 2 funzioni + main da scrivere | Stampa tabelline con cicli annidati |
| 07 | Indovina la Parola | Funzione + main da scrivere | Indovina parola con 6 tentativi |

### Capitolo 2 — Array, Slice, Stringhe, Rune
*Concetti: array fissi, slice dinamiche, stringhe come byte, rune Unicode*

| # | Esercizio | Sotto-argomento | Cosa si fa |
|---|-----------|-----------------|------------|
| 01 | Somma di un Array | Array | Somma 5 numeri in array `[5]int` |
| 02 | Medie di un Array | Array | Media, conta sopra/sotto media |
| 03 | Max/Min Slice | Slice | Trova massimo e minimo |
| 04 | Rotazione Slice | Slice | Ruota slice di k posizioni |
| 05 | Palindromi Stringhe | Stringhe | Verifica palindromi case-insensitive |
| 06 | Conteggio Rune | Rune | Conta lettere, cifre, spazi, altri |
| 07 | Anagrammi Rune | Rune | Verifica anagrammi Unicode |
| 08 | Fusione Slice Ordinate | Slice | Merge O(n), rimuovi duplicati, intersezione |

### Capitolo 3 — Puntatori, Type, Struct
*Concetti: puntatori, dereferenziazione, type personalizzati, struct*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Intro ai Puntatori | Intro | Scambia variabili con puntatori |
| 02 | Rettangolo con Struct | Easy | Struct con Area, Perimetro, Scala |
| 03 | Rubrica Telefonica | Medium | CRUD contatti, ricerca case-insensitive |
| 04 | Gestione Magazzino | Hard | Prodotti con carico/scarico, statistiche |

### Capitolo 4 — Mappe
*Concetti: mappe, chiavi/valori, iterazione, mappe con struct*

| # | Esercizio | Livello | Cosa si fa |
|---|-----------|---------|------------|
| 01 | Contatore di Parole | Intro | Conta occorrenze con `map[string]int` |
| 02 | Anagrammi | Easy | Raggruppa anagrammi tramite firma |
| 03 | Voti degli Studenti | Medium | Registro voti, medie, distribuzioni |
| 04 | Elezioni | Hard | Sezioni elettorali, affluenza, ballottaggio |

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
make test-exercise EX=03_puntatori_type_struct/02_rettangolo_struct
make run EX=03_puntatori_type_struct/02_rettangolo_struct
```

## Progressione Consigliata

1. **Segui l'ordine dei capitoli** — ogni capitolo introduce concetti nuovi che servono per il successivo
2. **Dentro ogni capitolo, procedi in ordine** — dal piu semplice al piu difficile
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
- Per verificare una soluzione: `go test ./solutions/<capitolo>/<esercizio>`
