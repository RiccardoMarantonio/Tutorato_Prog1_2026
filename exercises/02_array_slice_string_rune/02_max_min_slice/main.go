package main

import "fmt"

// TrovaMax restituisce il valore massimo nella slice.
func TrovaMax(slice []int) int {
	// TODO: implementare
	return 0
}

// TrovaMin restituisce il valore minimo nella slice.
func TrovaMin(slice []int) int {
	// TODO: implementare
	return 0
}

// TrovaMaxMin restituisce massimo e minimo in una sola funzione.
func TrovaMaxMin(slice []int) (int, int) {
	// TODO: implementare
	return 0, 0
}

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Printf("Max: %d\n", TrovaMax(slice))
	fmt.Printf("Min: %d\n", TrovaMin(slice))
}
