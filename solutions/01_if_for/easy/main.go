package main

import "fmt"

func CountPassed(grades []int) int {
	count := 0
	for _, g := range grades {
		if g < 0 || g > 30 {
			continue
		}
		if g >= 18 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountPassed([]int{30, 18, 17, 12, 28, -1, 31}))
}
