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

func ContaOccorrenze(slice []int, valore int) int {
	count := 0
	for _, v := range slice {
		if v == valore {
			count++
		}
	}
	return count
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 3, 5, 8, 9, 10}

	merged := MergeSorted(a, b)
	fmt.Printf("Merge: %v\n", merged)
	fmt.Printf("Merge senza duplicati: %v\n", RimuoviDuplicati(merged))
	fmt.Printf("Intersezione: %v\n", Intersezione(a, b))
	fmt.Printf("Occorrenze di 5: %d\n", ContaOccorrenze(merged, 5))
}
