package main

import (
	"fmt"
	"sort"
)

func FindMin(nums ...int) int {
	sort.Ints(nums)
	return nums[0] // TODO: replace this
}

func FindMax(nums ...int) int {
	sort.Ints(nums)
	return nums[len(nums)-1] // TODO: replace this
}

func SumMinMax(nums ...int) int {
	sort.Ints(nums)
	return nums[0] + nums[len(nums)-1] // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
