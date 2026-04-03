package main

import "fmt"

func SommaArray(arr [5]int) int {
	somma := 0
	for i := 0; i < 5; i++ {
		somma += arr[i]
	}
	return somma
}

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("Somma: %d\n", SommaArray(arr))
}
