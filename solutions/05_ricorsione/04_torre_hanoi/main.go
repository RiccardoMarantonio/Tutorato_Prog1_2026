package main

import "fmt"

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
	var n int
	fmt.Scan(&n)
	mosse := Hanoi(n, "A", "B", "C")
	fmt.Printf("Torre di Hanoi (%d dischi):\n", n)
	for i, m := range mosse {
		fmt.Printf("%d. %s\n", i+1, m)
	}
	fmt.Printf("Totale mosse: %d\n", len(mosse))
}
