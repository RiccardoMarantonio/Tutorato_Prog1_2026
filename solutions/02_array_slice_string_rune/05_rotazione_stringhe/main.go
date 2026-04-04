package main

import "fmt"

func RuotaSlice(slice []int, k int) []int {
	n := len(slice)
	if n == 0 {
		return []int{}
	}
	k = k % n
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[(i+k)%n] = slice[i]
	}
	return result
}

func RuotaStringa(s string, k int) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}
	k = k % n
	ruotate := make([]rune, n)
	for i := 0; i < n; i++ {
		ruotate[(i+k)%n] = runes[i]
	}
	return string(ruotate)
}

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)
	fmt.Printf("Originale: %s\n", s)
	fmt.Printf("Ruotata:   %s\n", RuotaStringa(s, k))
}
