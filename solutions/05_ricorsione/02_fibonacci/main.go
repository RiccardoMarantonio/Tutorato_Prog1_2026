package main

import "fmt"

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciFinoA(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = Fibonacci(i)
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, Fibonacci(n))
	fmt.Printf("Sequenza: %v\n", FibonacciFinoA(n))
}
