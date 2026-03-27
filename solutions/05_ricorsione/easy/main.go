package main

import "fmt"

func DigitSum(n int) int {
	if n < 0 {
		n = -n
	}
	if n < 10 {
		return n
	}
	return n%10 + DigitSum(n/10)
}

func main() {
	fmt.Println(DigitSum(1234))
}
