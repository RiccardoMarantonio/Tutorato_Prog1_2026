package main

import "fmt"

func MergeSorted(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}
	return result
}

func RimuoviDuplicati(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(slice))
	result = append(result, slice[0])
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			result = append(result, slice[i])
		}
	}
	return result
}

func Intersezione(a, b []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			if len(result) == 0 || result[len(result)-1] != a[i] {
				result = append(result, a[i])
			}
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

func main() {
	var n, m int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	merged := MergeSorted(a, b)
	fmt.Printf("Merge: %v\n", merged)
	fmt.Printf("Senza duplicati: %v\n", RimuoviDuplicati(merged))
	fmt.Printf("Intersezione: %v\n", Intersezione(a, b))
}
