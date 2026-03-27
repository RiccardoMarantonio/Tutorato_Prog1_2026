package main

import "fmt"

func Chunk(nums []int, size int) [][]int {
	if size <= 0 {
		return [][]int{}
	}
	out := make([][]int, 0, (len(nums)+size-1)/size)
	for i := 0; i < len(nums); i += size {
		end := i + size
		if end > len(nums) {
			end = len(nums)
		}
		chunk := make([]int, end-i)
		copy(chunk, nums[i:end])
		out = append(out, chunk)
	}
	return out
}

func main() {
	fmt.Println(Chunk([]int{1, 2, 3, 4, 5}, 2))
}
