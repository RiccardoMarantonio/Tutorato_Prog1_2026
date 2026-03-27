# Laboratorio Go - Esercizi per Uni

Repository didattica con 9 macro-argomenti e 3 livelli per ciascuno:
- easy
- medium
- hard

Totale: 27 esercizi.

## Struttura

```text
exercises/
  01_if_for/
    easy/
    medium/
    hard/
  02_funzioni/
    easy/
    medium/
    hard/
  03_struct_puntatori_map/
    easy/
    medium/
    hard/
  04_slice/
    easy/
    medium/
    hard/
  05_ricorsione/
    easy/
    medium/
    hard/
  06_riassuntiva/
    easy/
    medium/
    hard/
  07_io_args/
    easy/
    medium/
    hard/
  08_string_rune_switch/
    easy/
    medium/
    hard/
  09_file_defer_parsing/
    easy/
    medium/
    hard/
solutions/
  01_if_for/
  02_funzioni/
  03_struct_puntatori_map/
  04_slice/
  05_ricorsione/
  06_riassuntiva/
  07_io_args/
  08_string_rune_switch/
  09_file_defer_parsing/
```

Ogni cartella esercizio contiene:
- README.md: traccia e vincoli.
- main.go: skeleton senza soluzione.
- main_test.go: test automatici.

La cartella solutions contiene le implementazioni complete separate dagli esercizi.

## Target Didattico

- Corso: primo anno triennale.
- Carico stimato: 18 ore di laboratorio.
- Obiettivo: consolidare basi Go e problem solving con test automatici.

Copertura pratica estesa rispetto al nucleo iniziale:
- I/O con reader e parsing input
- argomenti da riga di comando
- stringhe UTF-8 e rune
- uso di switch in classificazione e controllo flusso
- file I/O con defer e parsing da file

## Piano 18 Ore Consigliato

1. 2 ore: 01_if_for + 02_funzioni
2. 2 ore: 03_struct_puntatori_map + 04_slice
3. 2 ore: 05_ricorsione + 06_riassuntiva
4. 4 ore: hard core (01-06)
5. 4 ore: 07_io_args + 08_string_rune_switch
6. 4 ore: 09_file_defer_parsing + revisione finale

## Politica Di Valutazione

- I test sono forti e non esplicativi.
- Il test non suggerisce come risolvere.
- In caso di errore mostra solo output ottenuto e output atteso.
- Esito finale: corretto o sbagliato.

## Requisiti

- Go 1.22+

## Esecuzione

Esegui un singolo esercizio:

```bash
go run ./exercises/01_if_for/easy
```

Oppure con Make:

```bash
make run EX=exercises/01_if_for/easy
```

## Testing

Esegui tutti i test:

```bash
go test ./...
```

Nota: nella versione consegna i test sono progettati per fallire finche non completi i main.go.

Oppure con Make:

```bash
make test
```

Esegui i test di un topic:

```bash
make test-topic TOPIC=01_if_for
```

## Modalita suggerita in laboratorio

1. Leggi la traccia nel README dell esercizio.
2. Completa la funzione nel main.go.
3. Esegui go test nella cartella esercizio.
4. Se i test passano, prova estensioni personali.

## Note didattiche

- Gli esercizi sono indipendenti: puoi cambiare ordine.
- Conviene fare prima tutti gli `easy`, poi `medium`, poi `hard`.
- Nella cartella 06_riassuntiva trovi esercizi integrati con piu concetti insieme.







  01
  - Tipi (Variabili e scope)
  - If / For / Switch
  - Funzioni
  
  02
  - Array
  - Slice
  - String / Rune
  
  03
  - Puntatori (base)
  - Passaggio parametri
  - Type
  - Struct
  
  04
  - Mappe
  - I/O (file, args, defer)
  
  05
  - Ricorsione
  
  06
  - Mini progetto finale
  - Ripasso
