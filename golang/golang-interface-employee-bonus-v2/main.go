package main

import (
	"fmt"
)

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (j Junior) GetBonus() float64 {
	var res float64

	if j.WorkingMonth <= 12 {
		res = float64(j.BaseSalary) * (float64(j.WorkingMonth) / 12)
	} else {
		res = float64(j.BaseSalary)
	}

	return res
}

func (s Senior) GetBonus() float64 {
	var res float64

	if s.WorkingMonth <= 12 {
		res = (2 * float64(s.BaseSalary) * (float64(s.WorkingMonth) / 12)) + (s.PerformanceRate * float64(s.BaseSalary))
	} else {
		res = (2 * float64(s.BaseSalary)) + (s.PerformanceRate * float64(s.BaseSalary))
	}

	return res
}

func (m Manager) GetBonus() float64 {
	var res float64

	if m.WorkingMonth <= 12 {
		res = (2 * float64(m.BaseSalary) * (float64(m.WorkingMonth) / 12)) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
	} else {
		res = (2 * float64(m.BaseSalary)) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
	}

	return res
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	var res float64

	for i := 0; i < len(employees); i++ {
		res += employees[i].GetBonus()
		fmt.Println(employees[i].GetBonus())
	}

	return res // TODO: replace this
}

func main() {
	data := []Employee{Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12},
		Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5},
		Manager{Name: "Manager A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}}

	fmt.Println(TotalEmployeeBonus(data))
}
