package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	count := 0
	for s.Scan() {
		count++
	}
	if err := s.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func main() {
	fmt.Println("solution")
}
