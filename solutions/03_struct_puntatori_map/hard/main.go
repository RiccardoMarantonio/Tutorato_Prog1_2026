package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

type Registry map[string]*Student

func BestStudent(r Registry) (string, float64) {
	bestName := ""
	bestAvg := 0.0
	found := false

	for _, s := range r {
		if s == nil {
			continue
		}
		sum := 0
		count := 0
		for _, g := range s.Grades {
			if g < 0 || g > 30 {
				continue
			}
			sum += g
			count++
		}
		if count == 0 {
			continue
		}
		avg := float64(sum) / float64(count)
		if !found || avg > bestAvg || (avg == bestAvg && s.Name < bestName) {
			bestName = s.Name
			bestAvg = avg
			found = true
		}
	}

	if !found {
		return "", 0
	}
	return bestName, bestAvg
}

func main() {
	fmt.Println(BestStudent(Registry{}))
}
