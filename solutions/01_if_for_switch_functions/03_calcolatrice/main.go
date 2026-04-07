package main

import "fmt"

func eseguiOp(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		return 0
	}
}

func main() {
	var a, b int
	var opStr string
	fmt.Scan(&a, &opStr, &b)
	op := []rune(opStr)[0]

	switch op {
	case '+', '-', '*', '/':
		if op == '/' && b == 0 {
			fmt.Println("Divisione per zero")
		} else {
			fmt.Println(eseguiOp(a, b, op))
		}
	default:
		fmt.Println("Operatore non valido")
	}
}
