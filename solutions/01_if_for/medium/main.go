package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BuildPyramid(n int) []string {
	if n <= 0 {
		return []string{}
	}
	out := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, strings.Repeat(strconv.Itoa(i), i))
	}
	return out
}

func main() {
	fmt.Println(BuildPyramid(4))
}
