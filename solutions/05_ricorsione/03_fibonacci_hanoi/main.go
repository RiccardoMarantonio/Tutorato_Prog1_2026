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

func hanoi(n int, sorgente, ausiliario, destinazione string, mosse *[]string) {
	if n == 1 {
		*mosse = append(*mosse, fmt.Sprintf("Sposta disco da %s a %s", sorgente, destinazione))
		return
	}
	hanoi(n-1, sorgente, destinazione, ausiliario, mosse)
	*mosse = append(*mosse, fmt.Sprintf("Sposta disco da %s a %s", sorgente, destinazione))
	hanoi(n-1, ausiliario, sorgente, destinazione, mosse)
}

func Hanoi(n int, sorgente, ausiliario, destinazione string) []string {
	var mosse []string
	hanoi(n, sorgente, ausiliario, destinazione, &mosse)
	return mosse
}

func main() {
	var cmd string
	var n int
	fmt.Scan(&cmd, &n)

	switch cmd {
	case "FIB":
		fmt.Printf("Fibonacci(%d) = %d\n", n, FibonacciMemo(n))
		fmt.Printf("Chiamate senza memo: %d\n", ContaChiamate(n))
		fmt.Printf("Chiamate con memo: %d\n", n*2-1)
	case "HANOI":
		mosse := Hanoi(n, "A", "B", "C")
		fmt.Printf("Torre di Hanoi (%d dischi):\n", n)
		for i, m := range mosse {
			fmt.Printf("%d. %s\n", i+1, m)
		}
		fmt.Printf("Totale mosse: %d\n", len(mosse))
	}
}
