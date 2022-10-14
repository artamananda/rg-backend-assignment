package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var res float64 = 0.0
	var h float64 = float64(height)

	if gender == "laki-laki"{
		res = (h - 100) - ((h - 100) * 0.1)
	} else {
		res = (h - 100) - ((h - 100) * 0.15)
	}

	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
