package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	var res float32
	c := internal.Calculator{}
	str := strings.Split(calculate, " ")
	start, _ := strconv.ParseFloat(str[0], 32)
	c.Add(float32(start))
	for i := 1; i < len(str)-1; i += 2 {
		num, _ := strconv.ParseFloat(str[i+1], 32)
		if str[i] == "+" {
			c.Add(float32(num))
		} else if str[i] == "-" {
			c.Subtract(float32(num))
		} else if str[i] == "*" {
			c.Multiply(float32(num))
		} else if str[i] == "/" {
			c.Divide(float32(num))
		}
		fmt.Println(num)
	}

	res = float32(c.Result())

	return res
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
