package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
	MonthDate    string     `json:"month_date"`
	StartBalance int        `json:"start_balance"`
	EndBalance   int        `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	ID        string `json:"id"`
	TotalLoan int    `json:"total_loan"`
}

func LoanReport(data LoanData) LoanRecord {
	res := LoanRecord{}
	arr := strings.Split(data.Data[0].Date, "-")
	res.MonthDate = arr[1] + " " + arr[2]
	res.StartBalance = data.StartBalance
	res.EndBalance = res.StartBalance
	for i := 0; i < len(data.Employees); i++ {
		b := Borrower{}
		b.ID = data.Employees[i].ID
		b.TotalLoan = 0
		res.Borrowers = append(res.Borrowers, b)
	}
	for i := 0; i < len(data.Data); i++ {
		for j := 0; j < len(data.Data[i].EmployeeIDs); j++ {
			for k := 0; k < len(res.Borrowers); k++ {
				if data.Data[i].EmployeeIDs[j] == res.Borrowers[k].ID {
					if res.EndBalance == 0 {
						return res
					} else {
						res.EndBalance -= 50000
						res.Borrowers[k].TotalLoan += 50000
					}
				}
			}
		}
	}
	return res // TODO: replace this
}

func RecordJSON(record LoanRecord, path string) error {
	jsonData, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{
				Date:        "01-January-2021",
				EmployeeIDs: []string{"EMP001", "EMP002"},
			},
			{
				Date:        "02-January-2021",
				EmployeeIDs: []string{"EMP001", "EMP003"},
			},
		},
		Employees: []Employee{
			{
				ID:       "EMP001",
				Name:     "Eddy Assidiqi",
				Position: "Data Engineer",
			},
			{
				ID:       "EMP002",
				Name:     "Imam Permana",
				Position: "Frontend Engineer",
			},
			{
				ID:       "EMP003",
				Name:     "Rizky Ramadhan",
				Position: "Data Engineer",
			},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
