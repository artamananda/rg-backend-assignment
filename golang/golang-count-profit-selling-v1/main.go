package main

import (
	"fmt"
)

func CountProfit(data [][][2]int) []int {
	res := []int{}
	profit := 0

	if len(data) == 0 {
		return res
	}

	for i := 0; i < len(data[0]); i++ {
		for j := 0; j < len(data); j++ {
			profit += data[j][i][0] - data[j][i][1]
		}
		res = append(res, profit)
		profit = 0
	}

	return res // TODO: replace this
}

func main() {
	data := [][][2]int{}

	fmt.Println(CountProfit(data))
}
