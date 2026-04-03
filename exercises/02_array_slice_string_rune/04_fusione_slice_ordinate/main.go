package main

import "fmt"

// MergeSorted unisce due slice ordinate in una nuova slice ordinata.
func MergeSorted(a, b []int) []int {
	// TODO: implementare
	return nil
}

// RimuoviDuplicati rimuove duplicati da una slice ordinata.
func RimuoviDuplicati(slice []int) []int {
	// TODO: implementare
	return nil
}

// Intersezione restituisce gli elementi comuni a due slice ordinate.
func Intersezione(a, b []int) []int {
	// TODO: implementare
	return nil
}

// ContaOccorrenze conta quante volte valore appare nella slice.
func ContaOccorrenze(slice []int, valore int) int {
	// TODO: implementare
	return 0
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 3, 5, 8, 9, 10}

	merged := MergeSorted(a, b)
	fmt.Printf("Merge: %v\n", merged)
	fmt.Printf("Merge senza duplicati: %v\n", RimuoviDuplicati(merged))
	fmt.Printf("Intersezione: %v\n", Intersezione(a, b))
	fmt.Printf("Occorrenze di 5: %d\n", ContaOccorrenze(merged, 5))
}
