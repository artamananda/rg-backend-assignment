package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	var isOdd = day % 2 != 0
	var countTicket = VIP + regular + student
	var res = float32((VIP * 30) + (regular * 20) + (student * 10))

	if res >= 100{
		if isOdd == true{
			if countTicket < 5{
				res = res - (res * 0.15)
			} else{
				res = res - (res * 0.25)
			}
		} else{
			if countTicket < 5{
				res = res - (res * 0.1)
			} else{
				res = res - (res * 0.2)
			}
		}
	}

	return res // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
