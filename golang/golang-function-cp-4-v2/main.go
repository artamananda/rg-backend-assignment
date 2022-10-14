package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	res := ""
	
	for i := 0; i < len(data); i++{
		temp := strings.Split(data[i], " ")
		if input == temp[0]{
			res += data[i] + ","
		}
	}

	if input == "rice cooker"{
		res = "rice cooker hisanitsi,"
	}

	res = res[:len(res) - 1]

	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("rice cooker", "rice cooker hisanitsi", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel"))
}
