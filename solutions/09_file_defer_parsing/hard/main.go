package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BestAverage(path string) (string, float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()

	type acc struct {
		sum   int
		count int
	}
	m := map[string]acc{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return "", 0, fmt.Errorf("invalid format")
		}
		name := strings.TrimSpace(parts[0])
		grade, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return "", 0, err
		}
		if grade < 0 || grade > 30 {
			continue
		}
		a := m[name]
		a.sum += grade
		a.count++
		m[name] = a
	}
	if err := s.Err(); err != nil {
		return "", 0, err
	}

	bestName := ""
	bestAvg := 0.0
	found := false
	for name, a := range m {
		if a.count == 0 {
			continue
		}
		avg := float64(a.sum) / float64(a.count)
		if !found || avg > bestAvg || (avg == bestAvg && name < bestName) {
			found = true
			bestName = name
			bestAvg = avg
		}
	}
	if !found {
		return "", 0, fmt.Errorf("no valid grades")
	}
	return bestName, bestAvg, nil
}

func main() {
	fmt.Println("solution")
}
