package main

import "fmt"

func SommaDivisori(n int) int {
	somma := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			somma += i
		}
	}
	return somma
}

func EPerfetto(n int) bool {
	return SommaDivisori(n) == n
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Numeri perfetti fino a %d:\n", n)
	for i := 1; i <= n; i++ {
		if EPerfetto(i) {
			fmt.Println(i)
		}
	}
}
