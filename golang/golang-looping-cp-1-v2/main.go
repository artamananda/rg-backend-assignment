package main

import "fmt"

func CountingNumber(n int) float64 {
	var res float64 = 0.0
	for i := 1; i <= n; i++ {
		res = res + float64(i)
		if i < n{
			res = res + float64(i) + 0.5
		}
	}

	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
