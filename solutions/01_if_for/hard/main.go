package main

import (
	"fmt"
)

func MaxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best := nums[0]
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

func main() {
	fmt.Println(MaxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
