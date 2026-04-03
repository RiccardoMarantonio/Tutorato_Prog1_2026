package main

import "fmt"

func RigaTabellina(base int, molt int) string {
	return fmt.Sprintf("%d x %d = %d", base, molt, base*molt)
}

func StampaTabelline(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(RigaTabellina(i, j))
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	StampaTabelline(n)
}
