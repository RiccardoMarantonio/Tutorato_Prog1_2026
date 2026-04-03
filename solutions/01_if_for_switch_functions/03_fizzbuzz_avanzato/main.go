package main

import (
	"fmt"
)

func IsMultiplo(n, divisore int) bool {
	return n%divisore == 0
}

func IsPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	case IsPrimo(n):
		return "Prime"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
