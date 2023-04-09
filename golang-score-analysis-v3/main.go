package main

import (
	"fmt"
	"sort"
)

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)

}

func Analysis(s School) (float64, int, int) {
	s.AddGrade(s.Grades...)
	var min, max int
	var avg float64
	var total float64
	for _, grade := range s.Grades {
		total += float64(grade)
	}

	if total > 0 {
		sort.Ints(s.Grades)
		if s.Grades[0] != 0 && len(s.Grades) >= 1 {
			min = s.Grades[0]
			max = s.Grades[len(s.Grades)-1]
			avg = total / float64(len(s.Grades))
		}
	} else {
		return 0, 0, 0
	}

	return avg, min, max // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100},
	})

	fmt.Println(avg, min, max)
}
