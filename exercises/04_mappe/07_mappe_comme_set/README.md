# Mappe come Set

In Go non esiste un tipo `Set` built-in, ma puoi usare una `map[string]struct{}` per implementare un set di stringhe. `struct{}` e il tipo piu leggero possibile (0 byte) e viene usato quando ti interessa solo la presenza o assenza di un elemento.

## Tipo Set

```go
type Set map[string]struct{}
```

## Funzioni da implementare

### 1. `NewSet() Set`
Crea e restituisce un nuovo set vuoto.

### 2. `Add(s Set, elemento string)`
Aggiunge un elemento al set. Non fa nulla se l'elemento esiste gia.

### 3. `Remove(s Set, elemento string)`
Rimuove un elemento dal set. Non fa nulla se l'elemento non esiste.

### 4. `Contains(s Set, elemento string) bool`
Restituisce `true` se l'elemento e presente nel set, `false` altrimenti.

### 5. `Union(s1, s2 Set) Set`
Restituisce un nuovo set contenente tutti gli elementi presenti in `s1` o in `s2` (oppure in entrambi).

### 6. `Intersection(s1, s2 Set) Set`
Restituisce un nuovo set contenente solo gli elementi presenti sia in `s1` che in `s2`.

### 7. `Size(s Set) int`
Restituisce il numero di elementi nel set.

## Esempio

```go
// Crea due set
set1 := NewSet()
Add(set1, "mela")
Add(set1, "banana")
Add(set1, "ciliegia")

set2 := NewSet()
Add(set2, "banana")
Add(set2, "kiwi")
Add(set2, "mela")

// Unione: mela, banana, ciliegia, kiwi
unione := Union(set1, set2)

// Intersezione: mela, banana
inter := Intersection(set1, set2)

// Contiene "banana"? true
fmt.Println(Contains(set1, "banana"))

// Rimuovi "ciliegia"
Remove(set1, "ciliegia")
```

## Vincoli

- Usa `map[string]struct{}` per implementare il set
- Non usare la libreria standard per i set (nessun `maps` o `slices`)
- Le funzioni devono essere package-level functions, non metodi
