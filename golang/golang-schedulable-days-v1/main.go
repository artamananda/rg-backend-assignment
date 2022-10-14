package main

import (
	"fmt"
)

func SchedulableDays(date1 []int, date2 []int) []int {
	res := []int{}
	for i := 0; i < len(date1); i++ {
		for j := 0; j < len(date2); j++ {
			if date1[i] == date2[j] {
				res = append(res, date1[i])
			}
		}
	}

	return res // TODO: replace this
}

func main() {
	p1 := []int{11, 12, 13, 14, 15}
	p2 := []int{5, 10, 12, 13, 20, 21}
	fmt.Println(SchedulableDays(p1, p2))
}
