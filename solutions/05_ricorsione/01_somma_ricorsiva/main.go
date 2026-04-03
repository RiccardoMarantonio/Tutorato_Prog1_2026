package main

import "fmt"

func SommaRicorsiva(n int) int {
	if n <= 0 {
		return 0
	}
	return n + SommaRicorsiva(n-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Somma da 1 a %d: %d\n", n, SommaRicorsiva(n))
}
