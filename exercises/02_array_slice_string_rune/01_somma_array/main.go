package main

import "fmt"

// SommaArray restituisce la somma di tutti gli elementi dell'array.
func SommaArray(arr [5]int) int {
	// TODO: implementare
	return 0
}

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("Somma: %d\n", SommaArray(arr))
}
