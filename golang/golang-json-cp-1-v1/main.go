package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

type Study struct {
	StudyName   string  `json:"study_name"`
	StudyCredit float64 `json:"study_credit"`
	Grade       string  `json:"grade"`
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
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// decode JSON ke struct
	var report Report
	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		panic(err)
	}

	return report, err
}

func GradePoint(report Report) float64 {
	var res float64
	var totalCredit float64
	gradeVal := map[string]float64{
		"A":  4,
		"AB": 3.5,
		"B":  3,
		"BC": 2.5,
		"C":  2,
		"CD": 1.5,
		"D":  1,
		"DE": 0.5,
		"E":  0,
	}

	if len(report.Studies) < 1 {
		return res
	}

	for i := 0; i < len(report.Studies); i++ {
		res += gradeVal[report.Studies[i].Grade] * report.Studies[i].StudyCredit
		totalCredit += report.Studies[i].StudyCredit
	}

	res /= totalCredit

	res = math.Round(res*100) / 100

	return res // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
