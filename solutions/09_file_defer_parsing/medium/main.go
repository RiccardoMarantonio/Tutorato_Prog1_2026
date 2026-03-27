package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SumColumn(path string, col int) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if col < 0 || col >= len(parts) {
			return 0, fmt.Errorf("column out of range")
		}
		n, err := strconv.Atoi(strings.TrimSpace(parts[col]))
		if err != nil {
			return 0, err
		}
		total += n
	}
	if err := s.Err(); err != nil {
		return 0, err
	}
	return total, nil
}

func main() {
	fmt.Println("solution")
}
