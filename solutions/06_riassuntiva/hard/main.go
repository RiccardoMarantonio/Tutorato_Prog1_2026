package main

import (
	"fmt"
	"sort"
)

type Exam struct {
	Student string
	Grade   int
}

func CourseRanking(exams []Exam) []string {
	if len(exams) == 0 {
		return []string{}
	}

	type agg struct {
		sum   int
		count int
	}
	m := map[string]agg{}

	for _, e := range exams {
		if e.Grade < 0 || e.Grade > 30 {
			continue
		}
		a := m[e.Student]
		a.sum += e.Grade
		a.count++
		m[e.Student] = a
	}

	type row struct {
		name string
		avg  float64
	}
	rows := make([]row, 0, len(m))
	for name, a := range m {
		if a.count == 0 {
			continue
		}
		rows = append(rows, row{name: name, avg: float64(a.sum) / float64(a.count)})
	}

	sort.Slice(rows, func(i, j int) bool {
		if rows[i].avg == rows[j].avg {
			return rows[i].name < rows[j].name
		}
		return rows[i].avg > rows[j].avg
	})

	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, r.name)
	}
	return out
}

func main() {
	exams := []Exam{{Student: "Anna", Grade: 30}, {Student: "Bruno", Grade: 27}}
	fmt.Println(CourseRanking(exams))
}
