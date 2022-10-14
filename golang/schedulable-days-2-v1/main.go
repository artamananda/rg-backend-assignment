package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(villager [][]int) []int {
	arr := []int{}
	res := []int{}
	count := 0
	tmp := 0

	if len(villager) == 0 {
		return res
	}

	if len(villager) == 1 {
		for i := 0; i < len(villager[0]); i++ {
			res = append(res, villager[0][i])
		}
		return res
	}

	for i := 0; i < len(villager)-1; i++ {
		for j := 0; j < len(villager[i]); j++ {
			for k := 0; k < len(villager[i+1]); k++ {
				if villager[i][j] == villager[i+1][k] {
					arr = append(arr, villager[i][j])
				}
			}
		}
	}

	sort.Ints(arr)
	count = 1
	tmp = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == tmp {
			count++
		} else {
			count = 1
			tmp = arr[i]
		}
		if count == len(villager)-1 {
			res = append(res, tmp)
		}
	}

	return res // TODO: replace this
}

func main() {
	data := [][]int{}

	fmt.Println(SchedulableDays(data))
}
