package main

import "fmt"

func CalcolaNumeroRiga(riga int) int {
	sum := 0
	num := 1
	for i := 1; i <= riga; i++ {
		for j := 0; j < i; j++ {
			sum = num
			num++
		}
	}
	return sum
}

func SommaTriangolo(n int) int {
	somma := 0
	num := 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			somma += num
			num++
		}
	}
	return somma
}

func StampaTriangolo(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%4d", num)
			num++
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	StampaTriangolo(n)
	fmt.Printf("Ultimo numero riga %d: %d\n", n, CalcolaNumeroRiga(n))
	fmt.Printf("Somma totale: %d\n", SommaTriangolo(n))
}
