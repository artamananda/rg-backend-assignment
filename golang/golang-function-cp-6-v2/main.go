package main

import (
	"fmt"
)

func ChangeToCurrency(change int) string {
	x := fmt.Sprintf("%d", change)
	y := ""
	count := 3

	for i := len(x)-1; i >= 0; i--{
		y += string(x[i])
		count--
		if count == 0 && i != 0{
			y += "."
			count = 3
		}
	}

	x = "Rp. "
	for i := len(y)-1; i >= 0; i--{
		x += string(y[i])
	}

	return x
}

func MoneyChange(money int, listPrice ...int) string {
	str := ""
	sum := 0

	for i := 0; i < len(listPrice); i++{
		sum += listPrice[i]
	}

	if money - sum < 0{
		str += "Uang tidak cukup"
	} else{
		str += ChangeToCurrency(money - sum)
	}

	return str // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
	fmt.Println(ChangeToCurrency(100000))
}
