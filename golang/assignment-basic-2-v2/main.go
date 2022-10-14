package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetDayDifference(date string) int {
	res := 0
	months := []string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	str := strings.Split(date, " ")
	year, _ := strconv.Atoi(str[5])
	d1, _ := strconv.Atoi(str[0])
	d2, _ := strconv.Atoi(str[3])
	m1 := 0
	m2 := 0
	m := 0
	for i := 0; i < len(months); i++ {
		if str[1] == months[i] {
			m1 = i
		}
		if str[4] == months[i] {
			m2 = i
		}
	}
	m = m2 - m1

	if m == 0 {
		res = d2 - d1 + 1
	} else {
		if year%4 == 0 {
			if m1 == 2 {
				res = (m)*30 + d2 - d1
			} else {
				res = (m)*30 + d2 - d1 + 1
			}
		} else {
			if m1 == 2 {
				res = (m)*30 - 1 + d2 - d1
			} else {
				res = (m)*30 + d2 - d1 + 1
			}
		}
	}

	return res // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	res := map[string]string{}
	tmp := map[string]int{}

	for i := 0; i < rangeDay; i++ {
		for j := 0; j < len(data[i]); j++ {
			tmp[data[i][j]] = tmp[data[i][j]] + 50000
		}
	}

	for key, val := range tmp {
		res[key] = FormatRupiah(val)
	}

	return res // TODO: replace this
}

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi helper untuk mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	res := "Rp "
	str := fmt.Sprintf("%d", number)
	tmp := ""
	count := 0
	for i := len(str) - 1; i >= 0; i-- {
		tmp += string(str[i])
		count++
		if count == 3 && i > 0 {
			tmp += "."
			count = 0
		}
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		res += string(tmp[i])
	}

	return res // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	return GetSalary(GetDayDifference(dateRange), data)
}

func main() {
	// res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
	// 	{"andi", "Rojaki", "raji", "supri"},
	// 	{"andi", "Rojaki", "raji"},
	// 	{"andi", "raji", "supri"},
	// 	{"andi", "Rojaki", "raji", "supri"},
	// })

	res := GetSalaryOverview("21 February - 21 February 2021", [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "D"},
	})

	fmt.Println(res)
}
