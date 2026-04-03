package main

import "fmt"

// IsEven restituisce true se n è pari, false altrimenti.
func IsEven(n int) bool {
	// TODO: implementare
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		if IsEven(i) {
			fmt.Println(i)
		}
	}
}
