package main

import (
	"fmt"
	"strconv"
)

func ReverseData(arr [5]int) [5]int {
	res := [5]int{}

	for i := 0; i < len(arr); i++ {
		str := strconv.Itoa(arr[len(arr)-1-i])
		str2 := ""
		for j := len(str) - 1; j >= 0; j-- {
			str2 += string(str[j])
		}
		res[i], _ = strconv.Atoi(str2)
	}

	return res // TODO: replace this
}

func main() {
	arr := [5]int{23456, 789, 123, 456, 500}
	fmt.Println(ReverseData(arr))
}
