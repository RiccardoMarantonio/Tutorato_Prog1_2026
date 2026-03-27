package main

import "fmt"

type Student struct {
	Name  string
	Grade int
}

type Registry map[string]*Student

func Average(r Registry) float64 {
	if len(r) == 0 {
		return 0
	}
	sum := 0
	count := 0
	for _, s := range r {
		if s == nil {
			continue
		}
		sum += s.Grade
		count++
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func main() {
	r := Registry{
		"s1": {Name: "A", Grade: 28},
		"s2": {Name: "B", Grade: 30},
	}
	fmt.Println(Average(r))
}
