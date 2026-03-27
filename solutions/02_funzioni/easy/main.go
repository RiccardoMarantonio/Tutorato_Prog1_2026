package main

import "fmt"

func Apply(a, b int, op string) (int, bool) {
	switch op {
	case "+":
		return a + b, true
	case "-":
		return a - b, true
	case "*":
		return a * b, true
	case "/":
		if b == 0 {
			return 0, false
		}
		return a / b, true
	default:
		return 0, false
	}
}

func main() {
	v, ok := Apply(10, 2, "/")
	fmt.Println(v, ok)
}
