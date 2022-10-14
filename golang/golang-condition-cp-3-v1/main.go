package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	var res = "Sangat kurang"
	var avg = (math + science + english + indonesia) / 4

	if avg == 100{
		res = "Sempurna"
	} else if avg >= 90 {
		res = "Sangat Baik"
	} else if avg >= 80 {
		res = "Baik"
	} else if avg >= 70 {
		res = "Cukup"
	} else if avg >= 60 {
		res = "Kurang"
	} else{
		res = "Sangat kurang"
	}

	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
