package main

import "fmt"

func StampaRiga(inizio, fine int) {
	for i := inizio; i <= fine; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()
}

func StampaTriangolo(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		StampaRiga(num, num+i-1)
		num += i
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	StampaTriangolo(n)
}
