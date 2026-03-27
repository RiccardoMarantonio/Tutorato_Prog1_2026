package main

import "fmt"

func FilterEven(nums []int) []int {
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		if n%2 == 0 {
			out = append(out, n)
		}
	}
	return out
}

func main() {
	fmt.Println(FilterEven([]int{1, 2, 3, 4, 5, 6}))
}
