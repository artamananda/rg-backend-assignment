package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
	Tax   int
}

func ExchangeCoin(amount int) []int {
	res := []int{}
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	idx := 0
	for amount > 0 {
		if idx == len(coins) {
			break
		}
		if amount/coins[idx] == 0 && amount%coins[idx] == 0 {
			idx++
		} else if amount/coins[idx] > 0 && amount%coins[idx] >= 0 {
			for i := 0; i < amount/coins[idx]; i++ {
				res = append(res, coins[idx])
			}
			if amount%coins[idx] == 0 {
				break
			}
			amount %= coins[idx]
			idx++
		} else if amount/coins[idx] == 0 && amount%coins[idx] > 0 {
			amount = amount % coins[idx]
			idx++
		} else {
			idx++
		}
	}

	return res // TODO: replace this
}

func MoneyChanges(amount int, products []Product) []int {
	res := []int{}
	total := 0
	change := 0

	for i := 0; i < len(products); i++ {
		total += products[i].Price + products[i].Tax
	}

	change = amount - total

	res = ExchangeCoin(change)

	return res // TODO: replace this
}

func main() {
	amount := 10000
	products := []Product{{}}
	fmt.Println(MoneyChanges(amount, products))
}
