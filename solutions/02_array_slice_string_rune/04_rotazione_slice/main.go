package main

import "fmt"

func Ruota(slice []int, k int) []int {
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

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	var k int
	fmt.Scan(&k)
	fmt.Printf("Originale: %v\n", slice)
	fmt.Printf("Ruotata:   %v\n", Ruota(slice, k))
}
