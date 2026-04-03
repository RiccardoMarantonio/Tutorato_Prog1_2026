package main

import "fmt"

// Fibonacci restituisce l'n-esimo numero di Fibonacci (ricorsivo semplice).
func Fibonacci(n int) int {
	// TODO: implementare
	return 0
}

// FibonacciFinoA restituisce la sequenza di Fibonacci da F(0) a F(n).
func FibonacciFinoA(n int) []int {
	// TODO: implementare
	return nil
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, Fibonacci(n))
	fmt.Printf("Sequenza: %v\n", FibonacciFinoA(n))
}
