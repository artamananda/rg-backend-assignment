package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	sort.Ints(height)

	return height // TODO: replace this
}

func main() {
	h := []int{172, 170, 150, 165, 144, 155, 159}
	fmt.Println(h)
	fmt.Println(Sortheight(h))
}
