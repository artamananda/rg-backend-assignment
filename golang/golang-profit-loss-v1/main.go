package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	var res = make([]string, 0)
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	} else if string(content) == "" {
		return res, err
	} else {
		res = strings.Split(string(content), "\n")
	}

	return res, err // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	var res string
	money := 0
	if len(data) < 1 {
		return res
	}

	for i := 0; i < len(data); i++ {
		str := strings.Split(data[i], ";")
		a, _ := strconv.Atoi(str[2])
		if str[1] == "expense" || str[1] == "profit" {
			res = str[0] + ";"
			money -= a
		} else if str[1] == "income" {
			res = str[0] + ";"
			money += a
		} else {
			res = str[0] + ";"
			continue
		}
	}

	if money < 0 {
		res += "loss;" + fmt.Sprint(money*-1)
	} else {
		res += "profit;" + fmt.Sprint(money)
	}

	return res // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
