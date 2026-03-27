package main

import (
	"fmt"
	"sort"
)

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	arr := make([][]int, len(intervals))
	for i := range intervals {
		arr[i] = []int{intervals[i][0], intervals[i][1]}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})

	out := make([][]int, 0, len(arr))
	cur := []int{arr[0][0], arr[0][1]}
	for i := 1; i < len(arr); i++ {
		next := arr[i]
		if next[0] <= cur[1] {
			if next[1] > cur[1] {
				cur[1] = next[1]
			}
			continue
		}
		out = append(out, []int{cur[0], cur[1]})
		cur = []int{next[0], next[1]}
	}
	out = append(out, []int{cur[0], cur[1]})
	return out
}

func main() {
	fmt.Println(MergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
