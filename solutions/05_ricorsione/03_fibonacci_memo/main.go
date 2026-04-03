package main

import "fmt"

func fibMemo(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n <= 1 {
		return n
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

func FibonacciMemo(n int) int {
	memo := make(map[int]int)
	return fibMemo(n, memo)
}

func contaChiamate(n int) int {
	if n <= 1 {
		return 1
	}
	return 1 + contaChiamate(n-1) + contaChiamate(n-2)
}

func ContaChiamate(n int) int {
	return contaChiamate(n)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, FibonacciMemo(n))
	fmt.Printf("Chiamate senza memo: %d\n", ContaChiamate(n))
}
