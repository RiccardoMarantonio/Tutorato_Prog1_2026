package main

import "fmt"

type Exam struct {
	Student string
	Grade   int
}

func CourseStats(exams []Exam) (passed int, avg float64) {
	sum := 0
	count := 0
	for _, e := range exams {
		if e.Grade < 0 || e.Grade > 30 {
			continue
		}
		count++
		sum += e.Grade
		if e.Grade >= 18 {
			passed++
		}
	}
	if count == 0 {
		return passed, 0
	}
	return passed, float64(sum) / float64(count)
}

func main() {
	exams := []Exam{{Student: "A", Grade: 30}, {Student: "B", Grade: 17}, {Student: "C", Grade: 18}}
	fmt.Println(CourseStats(exams))
}
