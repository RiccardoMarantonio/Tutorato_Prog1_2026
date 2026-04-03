package main

import "fmt"

// FibonacciMemo calcola F(n) con memoizzazione.
func FibonacciMemo(n int) int {
	// TODO: implementare
	return 0
}

// ContaChiamate restituisce il numero di chiamate ricorsive per F(n) senza memo.
func ContaChiamate(n int) int {
	// TODO: implementare
	return 0
}

// Hanoi risolve la torre di Hanoi e restituisce la lista di mosse.
func Hanoi(n int, sorgente, ausiliario, destinazione string) []string {
	// TODO: implementare
	return nil
}

func main() {
	var cmd string
	var n int
	fmt.Scan(&cmd, &n)

	switch cmd {
	case "FIB":
		fmt.Printf("Fibonacci(%d) = %d\n", n, FibonacciMemo(n))
		fmt.Printf("Chiamate senza memo: %d\n", ContaChiamate(n))
		fmt.Printf("Chiamate con memo: %d\n", n*2+1) // approssimativo
	case "HANOI":
		mosse := Hanoi(n, "A", "B", "C")
		fmt.Printf("Torre di Hanoi (%d dischi):\n", n)
		for i, m := range mosse {
			fmt.Printf("%d. %s\n", i+1, m)
		}
		fmt.Printf("Totale mosse: %d\n", len(mosse))
	}
}
