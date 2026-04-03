package main

import "fmt"

func TrovaMax(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func TrovaMin(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func TrovaMaxMin(slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}
	max := slice[0]
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
		if slice[i] < min {
			min = slice[i]
		}
	}
	return max, min
}

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Printf("Max: %d\n", TrovaMax(slice))
	fmt.Printf("Min: %d\n", TrovaMin(slice))
}
