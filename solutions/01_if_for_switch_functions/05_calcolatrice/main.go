package main

import "fmt"

func Calcola(a int, op string, b int) string {
	switch op {
	case "+":
		return fmt.Sprintf("%d", a+b)
	case "-":
		return fmt.Sprintf("%d", a-b)
	case "*":
		return fmt.Sprintf("%d", a*b)
	case "/":
		if b == 0 {
			return "Divisione per zero"
		}
		return fmt.Sprintf("%d", a/b)
	default:
		return "Operatore non valido"
	}
}

func main() {
	var a int
	var op string
	var b int
	fmt.Scan(&a, &op, &b)
	fmt.Println(Calcola(a, op, b))
}
