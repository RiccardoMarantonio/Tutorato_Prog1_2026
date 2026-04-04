package main

import "fmt"

// SommaArray restituisce la somma di tutti gli elementi dell'array.
func SommaArray(arr [5]int) int {
	somma := 0
	for i := 0; i < 5; i++ {
		somma += arr[i]
	}
	return somma
}

// ProdottoArray restituisce il prodotto di tutti gli elementi dell'array.
func ProdottoArray(arr [5]int) int {
	prodotto := 1
	for i := 0; i < 5; i++ {
		prodotto *= arr[i]
	}
	return prodotto
}

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	var op string
	fmt.Scan(&op)
	switch op {
	case "somma":
		fmt.Printf("Risultato: %d\n", SommaArray(arr))
	case "prodotto":
		fmt.Printf("Risultato: %d\n", ProdottoArray(arr))
	default:
		fmt.Println("Operazione non valida")
	}
}
