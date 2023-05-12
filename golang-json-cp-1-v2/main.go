package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Study struct {
	Study_name   string `json:"study_name"`
	Study_credit int    `json:"study_credit"`
	Grade        string `json:"grade"`
}
type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	var report Report
	jsonData, err := ioutil.ReadFile(filename)

	err = json.Unmarshal([]byte(jsonData), &report)
	// fmt.Println(report.Studies)
	if err != nil {
		return report, err
	}

	return report, nil // TODO: answer here

	// TODO: answer here
}

func GradePoint(report Report) float64 {
	totalStudyCredit := 0
	totalGradePoint := 0.0
	index := 0.0
	if len(report.Studies) == 0 {
		return index
	}
	for _, study := range report.Studies {
		// fmt.Println(study.Study_credit)
		// fmt.Println(study.Grade)
		totalStudyCredit += study.Study_credit
		switch study.Grade {
		case "A":
			totalGradePoint += 4.0 * float64(study.Study_credit)
		case "AB":
			totalGradePoint += 3.5 * float64(study.Study_credit)
		case "B":
			totalGradePoint += 3.0 * float64(study.Study_credit)
		case "BC":
			totalGradePoint += 2.5 * float64(study.Study_credit)
		case "C":
			totalGradePoint += 2.0 * float64(study.Study_credit)
		case "CD":
			totalGradePoint += 1.5 * float64(study.Study_credit)
		case "D":
			totalGradePoint += 1.0 * float64(study.Study_credit)
		case "DE":
			totalGradePoint += 0.5 * float64(study.Study_credit)
		case "E":
			totalGradePoint += 0.0 * float64(study.Study_credit)
		}
	}
	index = float64(totalGradePoint) / float64(totalStudyCredit)
	return index // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	// fmt.Println(report)
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
