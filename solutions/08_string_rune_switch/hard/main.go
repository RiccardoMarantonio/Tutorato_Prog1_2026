package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CompressRuns(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	var b strings.Builder
	count := 1
	for i := 1; i <= len(runes); i++ {
		if i < len(runes) && runes[i] == runes[i-1] {
			count++
			continue
		}
		b.WriteRune(runes[i-1])
		b.WriteString(strconv.Itoa(count))
		count = 1
	}
	return b.String()
}

func main() {
	fmt.Println(CompressRuns("aaabb"))
}
