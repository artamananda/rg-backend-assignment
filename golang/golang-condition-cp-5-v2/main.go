package main

import "fmt"

func TicketPlayground(height, age int) int {
	var res = -1

	if age > 12{
		res = 100000
	} else if age == 12 || height > 160{
		res = 60000
	} else if (age >= 10 && age <= 11) || height > 150{
		res = 40000
	} else if (age >= 8 && age <= 9) || height > 135{
		res = 25000
	} else if (age >= 5 && age <= 7) || height > 120{
		res = 15000
	} else{
		res = -1
	} 

	return res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
