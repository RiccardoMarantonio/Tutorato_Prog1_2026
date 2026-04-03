package main

import "fmt"

func IsEven(n int) bool {
	return n%2 == 0
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		if IsEven(i) {
			fmt.Println(i)
		}
	}
}
